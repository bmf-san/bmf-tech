---
title: Datadog continuous testingについて調べたことメモ
description: "Datadog継続的テスト機能を解説。ノーコードE2E、自己修復テスト、クロスブラウザ対応、CircleCI・Github Actions連携で信頼できるテスト運用を実現します。"
slug: datadog-continuous-testing-notes
date: 2023-01-31T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - Datadog
translation_key: datadog-continuous-testing-notes
---


# 概要
Datadog continuous testingについて調べたことのメモ。

# Datadog continuous testingとは
- ノーコード
    - 画面ポチポチでテストを用意できる
- 自己修復型E2E
    - "回復力"のあるテスト
    - 誤検出を最小限に抑える
- クロスブラウザテスト対応
- 主要なインテグレーションをカバー
    - CircleCI、Github Actions、Jenkins etc...

E2Eの運用を楽にし、テストの信頼性を保つ仕組みを備えている、ようだ。

# 導入
Chrome拡張の[Datadog test recorder](https://chrome.google.com/webstore/detail/datadog-test-recorder/kkbncfpddhdmkfmalecgnphegacgejoa)が必要。

それ以外の準備は不要ですぐに利用開始できる。

# テスト作成
1. Browser Testを作成
2. テストケースをレコーディングする

レコーディングする際は、popupで開くほうが良さそう（レコーディング画面内のUIはiframeなので）。

# テスト設定
テスト実行に関するあらゆるオプションの設定。

Basic認証の対応ができたり、Cookie、リクエストヘッダの設定などもできる。

cf. [テストコンフィギュレーション](https://docs.datadoghq.com/ja/synthetics/browser_tests/?tab=%E3%83%AA%E3%82%AF%E3%82%A8%E3%82%B9%E3%83%88%E3%82%AA%E3%83%97%E3%82%B7%E3%83%A7%E3%83%B3#%E3%83%86%E3%82%B9%E3%83%88%E3%82%B3%E3%83%B3%E3%83%95%E3%82%A3%E3%82%AE%E3%83%A5%E3%83%AC%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3)

# テスト項目
Browser Testsで利用できるテスト項目についてそれぞれ見てみる。

cf. https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B

## Assertion
DOM周りをチェックできるアサーション。

- [要素のコンテンツをテストする](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E8%A6%81%E7%B4%A0%E3%81%AE%E3%82%B3%E3%83%B3%E3%83%86%E3%83%B3%E3%83%84%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B)
- [要素の属性をテストする](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E8%A6%81%E7%B4%A0%E3%81%AE%E5%B1%9E%E6%80%A7%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B)
- [ある要素が存在するかどうかをテストする](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%81%82%E3%82%8B%E8%A6%81%E7%B4%A0%E3%81%8C%E5%AD%98%E5%9C%A8%E3%81%99%E3%82%8B%E3%81%8B%E3%81%A9%E3%81%86%E3%81%8B%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B)
    - CSSまたはXPath 1.0を選択して要素の存在をテストできる

## Navigation
遷移系。

- [ページを更新する](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%9A%E3%83%BC%E3%82%B8%E3%82%92%E6%9B%B4%E6%96%B0%E3%81%99%E3%82%8B)
    - スーパーリロードとかなさそう。
        - スーパーリロードしたい場合はJavaScriptで対応する感じかな
- [メールに移動してリンクをクリックする](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%A1%E3%83%BC%E3%83%AB%E3%81%AB%E7%A7%BB%E5%8B%95%E3%81%97%E3%81%A6%E3%83%AA%E3%83%B3%E3%82%AF%E3%82%92%E3%82%AF%E3%83%AA%E3%83%83%E3%82%AF%E3%81%99%E3%82%8B)
- [特定のリンクをたどる](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E7%89%B9%E5%AE%9A%E3%81%AE%E3%83%AA%E3%83%B3%E3%82%AF%E3%82%92%E3%81%9F%E3%81%A9%E3%82%8B)

## Special Actions
UIに関する操作系。

- [ホバー](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%9B%E3%83%90%E3%83%BC)
- [キーの押下](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%AD%E3%83%BC%E3%81%AE%E6%8A%BC%E4%B8%8B)
- [スクロール](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B9%E3%82%AF%E3%83%AD%E3%83%BC%E3%83%AB)
- [待機](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E5%BE%85%E6%A9%9F)

## Variables
任意の変数が定義できる。

- [パターン](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [要素](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E8%A6%81%E7%B4%A0)
    - ビルトインとして、numeric、alphabetic、alphanumeric、date、timestampが用意されている
    - テスト結果のローカル変数値を難読化するオプションもある
- [JavaScript](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#javascript)
    - 定義した関数をロード・実行できる
    - 同期・非同期両方をサポート
- [グローバル変数](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B0%E3%83%AD%E3%83%BC%E3%83%90%E3%83%AB%E5%A4%89%E6%95%B0)
    - [Synthetic Monitoring Settings](https://docs.datadoghq.com/ja/synthetics/settings/?tab=specifyvalue)で定義されたグローバル変数を利用できる
- [グローバル変数-MFA](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B0%E3%83%AD%E3%83%BC%E3%83%90%E3%83%AB%E5%A4%89%E6%95%B0---mfa)
    - [Synthetic Monitoring Settings](https://docs.datadoghq.com/ja/synthetics/settings/?tab=specifyvalue)で定義されたMFAグローバル変数を利用できる
        - MFA用のTOTPのサポートがある
            - [ブラウザテストにおける多要素認証 (MFA) 用 TOTP](https://docs.datadoghq.com/ja/synthetics/guide/browser-tests-totp/)
- [Email](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#email)
    - メールアドレスの送受信確認や、メール内容の確認、リンククリック等のテストで利用できるメールアドレスを作成できる
    - テスト実行の度に一意のメールアドレス・メールボックスが生成される

## Subtest
既存のブラウザテストを別のブラウザテストの中でも再利用できる。
最大2階層まで入れ子にして再利用することができる。

## HTTP Request
ブラウザテスト内でHTTPリクエストを実行できる。

# 自己修復
実際どれくらいのまでの修復が可能なのか細かいところはわかってない。
UIの変更があった際に自動的に変更を検知して探索するとは書かれている。

 > If there is a UI change that modifies an element (e.g., moves it to another location), the test will automatically locate the element again based on the points of reference that were not affected by the change. Once the test completes successfully, Datadog will recompute, or “self-heal,” any broken locators with updated values. This ensures that your tests do not break because of a simple UI change and can automatically adapt to the evolution of your application’s UI.
In the next section, we’ll look at how you can fine-tune your test notifications to ensure that you are only notified of legitimate failures.

cf. https://www.datadoghq.com/ja/blog/test-maintenance-best-practices/

あくまで”単純なUIの変更”に限るので、あまり期待を高めないようにするのが良さそう。

> テストが正常に実行されると、ブラウザテストは壊れたロケータを更新された値で再計算 (または「自己修復」) し、単純な UI の更新でテストが壊れることがなく、テストがアプリケーションの UI に自動的に適応することを保証します。

cf.  https://docs.datadoghq.com/ja/synthetics/browser_tests/advanced_options/

# 並列化
Syntheticの設定から並列化を設定できる。

cf. https://docs.datadoghq.com/ja/continuous_testing/settings/

# ダッシュボード
## Explore
Synthetic MonitoringとContinuous Testingの結果（CI Batches）およびテスト実行結果（Test Runs）を検索できる。
cf. https://docs.datadoghq.com/ja/continuous_testing/explorer/?tab=cibatches

## Test coverage
> RUM から収集したブラウザデータと Synthetic ブラウザのテスト結果を使用して、RUM アプリケーションのテストカバレッジ全体に関する洞察を提供
cf. https://docs.datadoghq.com/ja/synthetics/test_coverage/

どこがテストされていないのかといったことが分析できそう。テストケースの網羅性の改善に役立ちそう。
カバレッジの推移を追っていく

# CIインテグレーション
Synthetic testsは各種CIと連携できる。

用意されているインテグレーション。
- Azure DevOps Extension
- CircleCI Orb
- Github Actions
- GitLab
- Jenkins
- NPM package

cf. https://docs.datadoghq.com/continuous_testing/cicd_integrations/

# コード化
cf. https://registry.terraform.io/providers/DataDog/datadog/latest/docs/resources/synthetics_test

DOMベースのテストになるのでコードからテストケースを起こすということは厳しいと思う。。。

# 通知
Syntheticsとしての管理下にあるので通知周りは特に懸念なし。

# コスト
1000回/$12（オンデマンドだと$18）
cf. https://www.datadoghq.com/ja/pricing/?product=continuous-testing#continuous-testing

安価だと思う。
1000回以下は無料？？

# シナリオ管理方法
シナリオが沢山用意されたときに破綻しないか？

テストシナリオの管理方針としては、
- Views
- tag
を使う形になりそう。他の管理UIはなさそう。

テストケース自体の管理方針としては、DRYにテストを作っていくことが推奨されており、**サブテスト**を積極的に利用したほうが良さそう。
cf. https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B5%E3%83%96%E3%83%86%E3%82%B9%E3%83%88

# 気になった点メモ
- Browser TestではRUMと同じ情報を取得できるので、デバッグの参考情報にもなりそう
- Browser Test作成時にtest frequencyを設定する必要があるが、0にはできないので、定期実行する前提になる？
    - PAUSEDにしておけば大丈夫そうかも
- 並列化は必須になる気がする
    - 有効になっていないので有効化する必要がありそう？
        - add onぽい
            - cf. https://www.datadoghq.com/ja/pricing/?product=continuous-testing#continuous-testing
- データのクリーンアップに気をつける
    - テスト実行でデータが生成されるケースの場合、テストが実行途中でfailした場合、データが中途半端に生成された状態になる。
        - クリーンアップまで意識してテストケースを作ったとしても、テストが失敗すると途中のデータが残る。
            - クリーンアップをどう担保するか？
- テストでカバーできないと思われること
  - テスト実行デバイスにおけるOSの選択
    - ブラウザの種類やブラウザサイズは選択できるが、OSはできなそう
- SMS認証は対応していない模様
- ネイティブアプリは未対応
- プッシュ通知の検証なども厳しい
- E2Eに限った話ではないが、変更が多い部分やテストが壊れやすい部分を観測できるようにしておく仕組みを整えるの大事だなと思った。ダッシュボードがあるのでそれが可視化できそう。

# 参考
- [www.datadoghq.com - continuous-testing](https://www.datadoghq.com/ja/product/continuous-testing/)
- [docs.datadoghq.com](https://docs.datadoghq.com/ja/continuous_testing/)
- [www.datadoghq.com - Use Datadog Continuous Testing to release with confidence](https://www.datadoghq.com/ja/blog/release-confidently-with-datadog-continuous-testing/)
- [docs.datadoghq.com - Continuous Testing and CI/CD](https://docs.datadoghq.com/ja/continuous_testing/cicd_integrations/)
- [www.datadoghq.com - Best practices for creating end-to-end tests](https://www.datadoghq.com/ja/blog/test-creation-best-practices/)
- [www.datadoghq.com - Best practices for continuous testing with Datadog](https://www.datadoghq.com/ja/blog/best-practices-datadog-continuous-testing/)

