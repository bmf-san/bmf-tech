// slug_map_generator reads a MySQL dump and generates a CSV of post IDs +
// Japanese titles + English slug candidates for human review.
//
// Usage:
//
//	go run main.go [-sql PATH] [-out PATH]
//
// Defaults:
//
//	-sql  ../../bmf-tech_2026-03-01.sql
//	-out  ../../tools/slug_map.csv
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type postRecord struct {
	id, categoryID, title, createdAt string
}

func main() {
	sqlPath := flag.String("sql", "../../bmf-tech_2026-03-01.sql", "MySQL dump path")
	outPath := flag.String("out", "../../tools/slug_map.csv", "output CSV path")
	flag.Parse()

	records, err := parsePosts(*sqlPath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	fmt.Printf("parsed %d post records\n", len(records))

	if err := writeCSV(*outPath, records); err != nil {
		log.Fatalf("write csv: %v", err)
	}
	fmt.Printf("wrote %s\n", *outPath)
}

func parsePosts(path string) ([]postRecord, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 128*1024*1024) // 128 MB for large longtext fields
	scanner.Buffer(buf, cap(buf))

	var records []postRecord
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "INSERT INTO `posts` VALUES") {
			continue
		}
		records = append(records, parseTuples(line)...)
	}
	return records, scanner.Err()
}

// parseTuples extracts all (id,admin_id,cat_id,'title',...) rows from a bulk INSERT line.
func parseTuples(line string) []postRecord {
	var records []postRecord
	vi := strings.Index(line, " VALUES ")
	if vi < 0 {
		return nil
	}
	pos := vi + len(" VALUES ")

	for pos < len(line) {
		if line[pos] != '(' {
			pos++
			continue
		}
		pos++ // consume '('

		id, p := readInt(line, pos)
		pos = commaSkip(line, p)
		_, p = readInt(line, pos) // admin_id
		pos = commaSkip(line, p)
		catID, p := readInt(line, pos)
		pos = commaSkip(line, p)

		var title string
		if pos < len(line) && line[pos] == '\'' {
			title, p = readString(line, pos+1)
			pos = p
		} else if strings.HasPrefix(line[pos:], "NULL") {
			pos += 4
		}
		title = strings.Map(func(r rune) rune {
			if r == '\n' || r == '\r' {
				return ' '
			}
			return r
		}, title)

		// Skip md_body, html_body, status (3 fields)
		for i := 0; i < 3; i++ {
			pos = commaSkip(line, pos)
			pos = skipField(line, pos)
		}
		// created_at
		pos = commaSkip(line, pos)
		var createdAt string
		if pos < len(line) && line[pos] == '\'' {
			createdAt, p = readString(line, pos+1)
			pos = p
		}

		records = append(records, postRecord{
			id:         id,
			categoryID: catID,
			title:      title,
			createdAt:  createdAt,
		})

		// Advance to end of tuple ')'
		for pos < len(line) && line[pos] != ')' {
			if line[pos] == '\'' {
				_, pos = readString(line, pos+1)
			} else {
				pos++
			}
		}
		if pos < len(line) {
			pos++ // ')'
		}
		if pos < len(line) && line[pos] == ',' {
			pos++ // inter-tuple comma
		}
	}
	return records
}

func readInt(s string, pos int) (string, int) {
	start := pos
	for pos < len(s) && s[pos] >= '0' && s[pos] <= '9' {
		pos++
	}
	return s[start:pos], pos
}

func commaSkip(s string, pos int) int {
	if pos < len(s) && s[pos] == ',' {
		return pos + 1
	}
	return pos
}

// readString parses a MySQL single-quoted string starting after the opening quote.
func readString(s string, pos int) (string, int) {
	var b strings.Builder
	for pos < len(s) {
		c := s[pos]
		if c == '\\' && pos+1 < len(s) {
			nc := s[pos+1]
			switch nc {
			case '\'':
				b.WriteByte('\'')
			case '\\':
				b.WriteByte('\\')
			case 'n':
				b.WriteByte('\n')
			case 'r':
				b.WriteByte('\r')
			case '"':
				b.WriteByte('"')
			default:
				b.WriteByte('\\')
				b.WriteByte(nc)
			}
			pos += 2
		} else if c == '\'' {
			if pos+1 < len(s) && s[pos+1] == '\'' {
				b.WriteByte('\'')
				pos += 2
			} else {
				return b.String(), pos + 1
			}
		} else {
			b.WriteByte(c)
			pos++
		}
	}
	return b.String(), pos
}

// skipField advances past a NULL or single-quoted string field.
func skipField(s string, pos int) int {
	if pos >= len(s) {
		return pos
	}
	if s[pos] == '\'' {
		_, p := readString(s, pos+1)
		return p
	}
	if strings.HasPrefix(s[pos:], "NULL") {
		return pos + 4
	}
	return pos
}

func writeCSV(path string, records []postRecord) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	_ = w.Write([]string{"id", "category_id", "title_ja", "slug", "created_at", "reviewed"})
	for _, r := range records {
		_ = w.Write([]string{r.id, r.categoryID, r.title, generateSlug(r.id, r.title), r.createdAt, "false"})
	}
	return nil
}

func generateSlug(id, title string) string {
	var words []string
	var cur strings.Builder
	flush := func() {
		if cur.Len() > 1 {
			words = append(words, cur.String())
		}
		cur.Reset()
	}
	for _, r := range title {
		switch {
		case r == '+':
			flush()
			words = append(words, "plus")
		case r <= 127 && (unicode.IsLetter(r) || unicode.IsDigit(r)):
			cur.WriteRune(unicode.ToLower(r))
		default:
			flush()
		}
	}
	flush()

	// deduplicate adjacent
	var deduped []string
	for _, w := range words {
		if len(deduped) == 0 || deduped[len(deduped)-1] != w {
			deduped = append(deduped, w)
		}
	}
	if len(deduped) > 6 {
		deduped = deduped[:6]
	}
	if len(deduped) == 0 {
		return "post-" + id
	}
	return strings.Join(deduped, "-")
}
