# Changelog

## 0.6.0 (2026-03-03)

Full Changelog: [v0.5.2...v0.6.0](https://github.com/CaseMark/casedev-cli/compare/v0.5.2...v0.6.0)

### Features

* **api:** api update ([9c3c3ca](https://github.com/CaseMark/casedev-cli/commit/9c3c3ca2cc84fed8d734ac8b876332b8c94c2b58))

## 0.5.2 (2026-03-03)

Full Changelog: [v0.5.1...v0.5.2](https://github.com/CaseMark/casedev-cli/compare/v0.5.1...v0.5.2)

### Chores

* update SDK settings ([e79e380](https://github.com/CaseMark/casedev-cli/commit/e79e38056f07c5a67d0f340fef97b829a99a2629))

## 0.5.1 (2026-03-03)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/CaseMark/casedev-cli/compare/v0.5.0...v0.5.1)

### Bug Fixes

* deduplicate CI Slack alerts at workflow level, not PR level ([#26](https://github.com/CaseMark/casedev-cli/issues/26)) ([686a825](https://github.com/CaseMark/casedev-cli/commit/686a825f6a39455c31763430203519bd1a9e42f3))

## 0.5.0 (2026-03-03)

Full Changelog: [v0.4.2...v0.5.0](https://github.com/CaseMark/casedev-cli/compare/v0.4.2...v0.5.0)

### Features

* **api:** api update ([ebfe266](https://github.com/CaseMark/casedev-cli/commit/ebfe266a1b37a444e5beea911f8098ce97893d5e))


### Chores

* update SDK settings ([80009a7](https://github.com/CaseMark/casedev-cli/commit/80009a7b0313730d8f6e2ca366fc66aed0b304b0))

## 0.4.2 (2026-02-27)

Full Changelog: [v0.4.1...v0.4.2](https://github.com/CaseMark/casedev-cli/compare/v0.4.1...v0.4.2)

### Bug Fixes

* more gracefully handle empty stdin input ([53c41d4](https://github.com/CaseMark/casedev-cli/commit/53c41d4fbaf3f47c90eb4b0e43c2aadc621c0801))

## 0.4.1 (2026-02-27)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/CaseMark/casedev-cli/compare/v0.4.0...v0.4.1)

### Chores

* zip READMEs as part of build artifact ([e825863](https://github.com/CaseMark/casedev-cli/commit/e825863966a1910c8761141fe7bfc82ed95ba0e7))

## 0.4.0 (2026-02-26)

Full Changelog: [v0.3.6...v0.4.0](https://github.com/CaseMark/casedev-cli/compare/v0.3.6...v0.4.0)

### Features

* **api:** api update ([7c7c979](https://github.com/CaseMark/casedev-cli/commit/7c7c9795e927b6d31077647494c4f9f11db86357))


### Bug Fixes

* bump casedev-go to v0.4.0 to fix CI build failures ([fafab21](https://github.com/CaseMark/casedev-cli/commit/fafab211a854923b7e041d6838d89332bbd352eb))

## 0.3.6 (2026-02-25)

Full Changelog: [v0.3.5...v0.3.6](https://github.com/CaseMark/casedev-cli/compare/v0.3.5...v0.3.6)

### Bug Fixes

* pin formatting for headers to always use repeat/dot formats ([a33f7f4](https://github.com/CaseMark/casedev-cli/commit/a33f7f4d269575b54b159c24ee0689b3f4ce134b))


### Chores

* **internal:** codegen related update ([2e8f723](https://github.com/CaseMark/casedev-cli/commit/2e8f72335f3fc89aa3821fcb826af32eb88bd974))
* update readme with better instructions for installing with homebrew ([90b772a](https://github.com/CaseMark/casedev-cli/commit/90b772a0393b3eba09af6e85334462f831990006))

## 0.3.5 (2026-02-24)

Full Changelog: [v0.3.4...v0.3.5](https://github.com/CaseMark/casedev-cli/compare/v0.3.4...v0.3.5)

### Chores

* update SDK settings ([5c433ee](https://github.com/CaseMark/casedev-cli/commit/5c433eef255aa1bdea5b63113a2f3b4bebe72c03))

## 0.3.4 (2026-02-24)

Full Changelog: [v0.3.3...v0.3.4](https://github.com/CaseMark/casedev-cli/compare/v0.3.3...v0.3.4)

### Bug Fixes

* prevent duplicate Slack alerts from release-please branches ([e72bd51](https://github.com/CaseMark/casedev-cli/commit/e72bd5166ad6d84573be5f76cabaffcb707da72d))

## 0.3.3 (2026-02-23)

Full Changelog: [v0.3.2...v0.3.3](https://github.com/CaseMark/casedev-cli/compare/v0.3.2...v0.3.3)

### Chores

* update SDK settings ([b0396d5](https://github.com/CaseMark/casedev-cli/commit/b0396d5e5f6d4525e1d8bb0969543616b9b5be92))

## 0.3.2 (2026-02-23)

Full Changelog: [v0.3.1...v0.3.2](https://github.com/CaseMark/casedev-cli/compare/v0.3.1...v0.3.2)

### Bug Fixes

* only alert on CI failure for PRs targeting main ([d342c99](https://github.com/CaseMark/casedev-cli/commit/d342c997e9a2c7671c45cda992286d105106a86e))
* only alert Slack on CI failure for main branch ([435b6d5](https://github.com/CaseMark/casedev-cli/commit/435b6d59171c04c26e8be50ab9842c4d86886b40))

## 0.3.1 (2026-02-23)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/CaseMark/casedev-cli/compare/v0.3.0...v0.3.1)

### Chores

* replace PR open Slack alert with CI failure alert ([b414a29](https://github.com/CaseMark/casedev-cli/commit/b414a29ba634fe75e5fdf2b8c33b4e32f9d55f23))

## 0.3.0 (2026-02-21)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/CaseMark/casedev-cli/compare/v0.2.1...v0.3.0)

### Features

* **api:** api update ([25967a9](https://github.com/CaseMark/casedev-cli/commit/25967a98543354995c77341905c1d05f8eca53cb))
* **api:** api update ([822ccc7](https://github.com/CaseMark/casedev-cli/commit/822ccc73492432fe86d11e7f8b2ed704d91a8d6d))


### Bug Fixes

* bump casedev-go to v0.3.0 for TrademarkSearch support ([f39a0e3](https://github.com/CaseMark/casedev-cli/commit/f39a0e3179e469c416e273d28dade87146c8bbff))
* set dummy CASEDEV_API_KEY for mock server tests ([9701a85](https://github.com/CaseMark/casedev-cli/commit/9701a85649f1d6c215a2caec5983b7c94d7089f6))

## 0.2.1 (2026-02-20)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/CaseMark/casedev-cli/compare/v0.2.0...v0.2.1)

## 0.2.0 (2026-02-20)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/CaseMark/casedev-cli/compare/v0.1.0...v0.2.0)

### Features

* **api:** api update ([0e0fd94](https://github.com/CaseMark/casedev-cli/commit/0e0fd946aad66fe7b61eab40f2f58786bc5f4bde))


### Chores

* **internal:** codegen related update ([a27a6d2](https://github.com/CaseMark/casedev-cli/commit/a27a6d23a0d78a1b09264c26a826ca9015a8f439))
* **internal:** remove mock server code ([6b046c7](https://github.com/CaseMark/casedev-cli/commit/6b046c779afbab408f53509c75d006e01d9bea96))
* update mock server docs ([5eefc0f](https://github.com/CaseMark/casedev-cli/commit/5eefc0ff576f9fc46026740219fd20f0198fd0f1))

## 0.1.0 (2026-02-19)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/CaseMark/casedev-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** api update ([f8e26f5](https://github.com/CaseMark/casedev-cli/commit/f8e26f51d21ca18956b62c140c783189907620e2))
