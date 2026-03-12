---
title: Service Downtime Due to File System Capacity Shortage
description: An in-depth look at Service Downtime Due to File System Capacity Shortage, covering key concepts and practical insights.
slug: service-outage-due-to-storage
date: 2023-05-07T00:00:00Z
author: bmf-san
categories:
  - Incident Report
tags:
  - Postmortem
translation_key: service-outage-due-to-storage
---



# Status
Resolved

# Incident
On January 2, 2023, around noon, it was discovered that accessing `https://bmf-tech.com/` resulted in slow responses and constant 500 errors. Attempts to log into Grafana for investigation failed.

Considering the possibility that some containers might have gone down for some reason, a deployment was carried out. However, upon confirming the error log `no space left on device`, another cause was suspected.

# Impact
All services of bmf-tech were unavailable.

Upon checking the Nginx request status, the service had been down since around 5:48 AM on January 2, 2023.

Recovery occurred around 12:40 PM on the same day.

![Screenshot 2023-01-02 13 26 53](/assets/images/posts/service-outage-due-to-storage/210195174-bf6c78c5-505d-41ff-8329-7781ef1fcae1.png)

Between 5:48 AM and 12:40 PM on January 2, 2023, 58 instances of 500 errors occurred.
*Although we wanted to measure a certain number of users, it was difficult to investigate as we had not adjusted to aggregate from logs or GA4, etc.*

![Screenshot 2023-01-02 13 48 34](/assets/images/posts/service-outage-due-to-storage/210195945-c9e1ae53-d624-4119-b8c5-22a185e66239.png)

# Cause
![Screenshot 2023-01-02 12 39 03](/assets/images/posts/service-outage-due-to-storage/210195173-9bc2975f-73f4-495d-bb19-71732ce593f2.png)

The cause was the lack of free space in the file system.

# Response
Deleted data that was consuming file system capacity to secure free space.

# Actions
1. SSH into the server and check disk free space with the `df` command. There was no free space.
```sh
df -h
```
2. Investigate the area using the most disk space with the `du` command. Identified as under `/var/lib/docker/`.
```sh
du -sh /var/lib/docker*
```
3. Delete unused Docker objects to secure free space.
```sh
docker system prune -a
```
4. Additionally, deleted journal logs that were consuming the second most space, leaving 200M.
```sh
journalctl --vacuum-size=200M
```

# Addendum
The above response was insufficient.

Log files accumulating under `/var/lib/docker/containers` were taking up significant space, so the container log rotation was adjusted to address this.

```yml
ex. 
  app:
    container_name: "app"
    logging:
      driver: "json-file"
      options:
        tag: "{{.ImageName}}|{{.Name}}|{{.ImageFullID}}|{{.FullID}}"
        max-size: "10m" // File size to rollover
        max-file: "3" // Number of rollovers before logs are discarded
```

This response allowed for significant free space. This was the main bottleneck...

# Prevention
Added file system usage to alerts to detect and address capacity issues in advance.

![Screenshot 2023-01-02 14 10 54](/assets/images/posts/service-outage-due-to-storage/210196950-64e37302-38a2-4cda-a34e-e20b99b77791.png)

# Others
We want to identify data that can be deleted or should be rotated to save space.

# Thoughts
This was the first incident since replacing bmf-tech.
