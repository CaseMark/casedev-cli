# Changelog

## 0.25.0 (2026-03-31)

Full Changelog: [v0.24.0...v0.25.0](https://github.com/CaseMark/casedev-cli/compare/v0.24.0...v0.25.0)

### Features

* **api:** api update ([6a331c2](https://github.com/CaseMark/casedev-cli/commit/6a331c2248b088fcc01fc7c12ac9b45531f35373))

## 0.24.0 (2026-03-31)

Full Changelog: [v0.23.0...v0.24.0](https://github.com/CaseMark/casedev-cli/compare/v0.23.0...v0.24.0)

### Features

* **api:** api update ([dc8a45c](https://github.com/CaseMark/casedev-cli/commit/dc8a45cd071c7640d8bacbf4b25cffd84b3acb30))

## 0.23.0 (2026-03-30)

Full Changelog: [v0.22.0...v0.23.0](https://github.com/CaseMark/casedev-cli/compare/v0.22.0...v0.23.0)

### Features

* add default description for enum CLI flags without an explicit description ([c564ecb](https://github.com/CaseMark/casedev-cli/commit/c564ecbaeafc59ae38bc5a1f6890d2450eec3eb0))
* **api:** api update ([d638c0b](https://github.com/CaseMark/casedev-cli/commit/d638c0b1e3d9279ab75408ce506d14d1c7a32c1d))
* **api:** api update ([30d3adf](https://github.com/CaseMark/casedev-cli/commit/30d3adf6b533ca86f32b503b26363d8ed002d71c))
* **api:** api update ([78eb231](https://github.com/CaseMark/casedev-cli/commit/78eb231eb2f916e5234c232dbc0ebf51e7ad24f0))
* set CLI flag constant values automatically where `x-stainless-const` is set ([9ed07b2](https://github.com/CaseMark/casedev-cli/commit/9ed07b218aca315d5b740c3695b0c77dce8c91c7))


### Bug Fixes

* cli no longer hangs when stdin is attached to a pipe with empty input ([136a9fb](https://github.com/CaseMark/casedev-cli/commit/136a9fb53c2f3e83054711fbc0c5d01718bf3ac5))
* fix for off-by-one error in pagination logic ([db09c87](https://github.com/CaseMark/casedev-cli/commit/db09c872956e4c7c038a6d4f0dadb5b60e44a910))


### Chores

* **ci:** skip lint on metadata-only changes ([baf033b](https://github.com/CaseMark/casedev-cli/commit/baf033bf1d0af73628e0c6877d364c22ddcdd33c))
* **internal:** codegen related update ([23e1822](https://github.com/CaseMark/casedev-cli/commit/23e18227417db0663835bd61037bff322337b8ad))
* **internal:** codegen related update ([79f80bc](https://github.com/CaseMark/casedev-cli/commit/79f80bceb5ce0ce97c0c9fe79e2aef74eb29d2e1))
* **internal:** update gitignore ([391e22e](https://github.com/CaseMark/casedev-cli/commit/391e22ee305269489ee8737a53e5e1b072015031))
* **internal:** update multipart form array serialization ([9680caf](https://github.com/CaseMark/casedev-cli/commit/9680caf8afbb66bfe8266ab9ccd4ada71b046530))
* omit full usage information when missing required CLI parameters ([aba147c](https://github.com/CaseMark/casedev-cli/commit/aba147ccffe743a48b2d3140b0787cecd2db100f))
* **tests:** bump steady to v0.19.4 ([8a8ce29](https://github.com/CaseMark/casedev-cli/commit/8a8ce29d99f5b155ba8b40e71a4d23376f23adaa))
* **tests:** bump steady to v0.19.5 ([1b3ef46](https://github.com/CaseMark/casedev-cli/commit/1b3ef46113bf79ca6426b4d49268bcabaac58bfe))
* **tests:** bump steady to v0.19.6 ([c51c368](https://github.com/CaseMark/casedev-cli/commit/c51c368560db70442f2a2b02f5eaddeacb57ec98))
* **tests:** bump steady to v0.19.7 ([2f8e158](https://github.com/CaseMark/casedev-cli/commit/2f8e1588667c4420c613ae60c5c1f982f17102be))


### Refactors

* **tests:** switch from prism to steady ([4f50ad5](https://github.com/CaseMark/casedev-cli/commit/4f50ad5efa6345b97bf9f69beaaa3669e393cb1f))

## 0.22.0 (2026-03-19)

Full Changelog: [v0.21.1...v0.22.0](https://github.com/CaseMark/casedev-cli/compare/v0.21.1...v0.22.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([161d0f2](https://github.com/CaseMark/casedev-cli/commit/161d0f2f88a18632f6b01edb2c1a10ad8f5186aa))
* add support for file downloads from binary response endpoints ([ca496e8](https://github.com/CaseMark/casedev-cli/commit/ca496e865692771ab3dfb1f99f97382e25ac3fc6))
* **api:** api update ([302e01a](https://github.com/CaseMark/casedev-cli/commit/302e01af3df9fe604bfcde198c03f1beaca17771))
* **api:** api update ([9abcbc3](https://github.com/CaseMark/casedev-cli/commit/9abcbc3ac32c890b8e800100ec1179a910634421))
* **api:** api update ([55398aa](https://github.com/CaseMark/casedev-cli/commit/55398aaee864618bcbbe0823c8e283475b060a5b))
* **api:** api update ([7d0fdb9](https://github.com/CaseMark/casedev-cli/commit/7d0fdb98e3e83a7c66da614d49e8f04b132490c1))
* **api:** api update ([e06c732](https://github.com/CaseMark/casedev-cli/commit/e06c7324a2defa3ff645a8a7d83cf6b30be22b20))
* **api:** api update ([f5c04cd](https://github.com/CaseMark/casedev-cli/commit/f5c04cdd99d24af359d313d85c5124b14bde4d8b))
* **api:** api update ([0673986](https://github.com/CaseMark/casedev-cli/commit/0673986c650e889d56636668b7b7434838b279f4))
* **api:** api update ([bbf3531](https://github.com/CaseMark/casedev-cli/commit/bbf3531a2bddc2d903a3ceb8c1bd9972220826a4))
* **api:** api update ([d530243](https://github.com/CaseMark/casedev-cli/commit/d5302438f10ebf09d244939e0a2f906fb7e75e3e))
* **api:** api update ([cc9a197](https://github.com/CaseMark/casedev-cli/commit/cc9a197be9209d744372909c690b222232b48949))
* **api:** api update ([57dc9a8](https://github.com/CaseMark/casedev-cli/commit/57dc9a8f52e1a0ffc0219d6d7cc69323d1fd57f7))
* **api:** api update ([c3b8f51](https://github.com/CaseMark/casedev-cli/commit/c3b8f5128945bc7eaff19e42659eeea0f84f7e5a))
* **api:** api update ([a99754f](https://github.com/CaseMark/casedev-cli/commit/a99754fd54fe7104c2eebba99af7b70318c7adab))
* **api:** api update ([fc663ad](https://github.com/CaseMark/casedev-cli/commit/fc663adddf3eca77a8e658de11de42eb4952c79d))
* **api:** api update ([9c3c3ca](https://github.com/CaseMark/casedev-cli/commit/9c3c3ca2cc84fed8d734ac8b876332b8c94c2b58))
* **api:** api update ([ebfe266](https://github.com/CaseMark/casedev-cli/commit/ebfe266a1b37a444e5beea911f8098ce97893d5e))
* **api:** api update ([7c7c979](https://github.com/CaseMark/casedev-cli/commit/7c7c9795e927b6d31077647494c4f9f11db86357))
* **api:** api update ([25967a9](https://github.com/CaseMark/casedev-cli/commit/25967a98543354995c77341905c1d05f8eca53cb))
* **api:** api update ([822ccc7](https://github.com/CaseMark/casedev-cli/commit/822ccc73492432fe86d11e7f8b2ed704d91a8d6d))
* **api:** api update ([0e0fd94](https://github.com/CaseMark/casedev-cli/commit/0e0fd946aad66fe7b61eab40f2f58786bc5f4bde))
* **api:** api update ([f8e26f5](https://github.com/CaseMark/casedev-cli/commit/f8e26f51d21ca18956b62c140c783189907620e2))
* improved documentation and flags for client options ([94e1cb9](https://github.com/CaseMark/casedev-cli/commit/94e1cb9ba8b19d54e2729ed89b9869659ff1e62f))
* support passing required body params through pipes ([e61b278](https://github.com/CaseMark/casedev-cli/commit/e61b27882931978ef068b50f759b2c22c0fa965d))


### Bug Fixes

* add casedev-go checksum entries ([#46](https://github.com/CaseMark/casedev-cli/issues/46)) ([11c0b42](https://github.com/CaseMark/casedev-cli/commit/11c0b424578eca073b5b73b54df9f4aa50fe74ee))
* align casedev-go version with generated agent APIs ([01b54d0](https://github.com/CaseMark/casedev-cli/commit/01b54d0f37daee74fada0d97079d1cc4c0e25ed8))
* avoid printing usage errors twice ([1f49bd9](https://github.com/CaseMark/casedev-cli/commit/1f49bd9c3a000f057be6ba88b400b2bd11dea74a))
* avoid reading from stdin unless request body is form encoded or json ([72e551b](https://github.com/CaseMark/casedev-cli/commit/72e551bee1115cc80c456f1aa60090837f064c8f))
* better support passing client args in any position ([1abc117](https://github.com/CaseMark/casedev-cli/commit/1abc1174fd9e3f4bfd069c139e88d485bd2666fd))
* bump casedev-go for chat file commands ([e4ce095](https://github.com/CaseMark/casedev-cli/commit/e4ce095152faa5fa95483066b34d7a60e5bf09d7))
* bump casedev-go from v0.8.0 to v0.9.0 ([#36](https://github.com/CaseMark/casedev-cli/issues/36)) ([e370249](https://github.com/CaseMark/casedev-cli/commit/e37024979013b8479d16ed24e36fdfad27bbd60f))
* bump casedev-go to v0.19.0 for CLI build ([7845550](https://github.com/CaseMark/casedev-cli/commit/7845550be0483eb08ab398da3d9b37074ffe653d))
* bump casedev-go to v0.3.0 for TrademarkSearch support ([f39a0e3](https://github.com/CaseMark/casedev-cli/commit/f39a0e3179e469c416e273d28dade87146c8bbff))
* bump casedev-go to v0.4.0 to fix CI build failures ([fafab21](https://github.com/CaseMark/casedev-cli/commit/fafab211a854923b7e041d6838d89332bbd352eb))
* clear quarantine for homebrew installs ([18af47f](https://github.com/CaseMark/casedev-cli/commit/18af47f8fb8aeefde6aefae39613637f450c65fa))
* deduplicate CI Slack alerts at workflow level, not PR level ([#26](https://github.com/CaseMark/casedev-cli/issues/26)) ([686a825](https://github.com/CaseMark/casedev-cli/commit/686a825f6a39455c31763430203519bd1a9e42f3))
* fix for encoding arrays with `any` type items ([66aca04](https://github.com/CaseMark/casedev-cli/commit/66aca04587d2e606a8c90ff5e10f0b64dea4b3b2))
* fix for test cases with newlines in YAML and better error reporting ([8735c70](https://github.com/CaseMark/casedev-cli/commit/8735c70b6a32e7abf9252c94a14d76f0a70204d3))
* improve linking behavior when developing on a branch not in the Go SDK ([5cae562](https://github.com/CaseMark/casedev-cli/commit/5cae5621bd05cdd2ed4501626a54d556e7106842))
* improved workflow for developing on branches ([561bc6e](https://github.com/CaseMark/casedev-cli/commit/561bc6e0671fb1b79d86a8e59f3847edac60fcf3))
* more gracefully handle empty stdin input ([53c41d4](https://github.com/CaseMark/casedev-cli/commit/53c41d4fbaf3f47c90eb4b0e43c2aadc621c0801))
* no longer require an API key when building on production repos ([27607fe](https://github.com/CaseMark/casedev-cli/commit/27607fecc4a04348461621afd8abf95ca74260d5))
* only alert on CI failure for PRs targeting main ([d342c99](https://github.com/CaseMark/casedev-cli/commit/d342c997e9a2c7671c45cda992286d105106a86e))
* only alert Slack on CI failure for main branch ([435b6d5](https://github.com/CaseMark/casedev-cli/commit/435b6d59171c04c26e8be50ab9842c4d86886b40))
* only set client options when the corresponding CLI flag or env var is explicitly set ([fd1ab05](https://github.com/CaseMark/casedev-cli/commit/fd1ab052cd8127b05b159aa2a1b0ab931cdd9d62))
* pin formatting for headers to always use repeat/dot formats ([a33f7f4](https://github.com/CaseMark/casedev-cli/commit/a33f7f4d269575b54b159c24ee0689b3f4ce134b))
* prevent duplicate Slack alerts from release-please branches ([e72bd51](https://github.com/CaseMark/casedev-cli/commit/e72bd5166ad6d84573be5f76cabaffcb707da72d))
* set dummy CASEDEV_API_KEY for mock server tests ([9701a85](https://github.com/CaseMark/casedev-cli/commit/9701a85649f1d6c215a2caec5983b7c94d7089f6))
* skip sumdb for sibling CaseMark modules ([e2839eb](https://github.com/CaseMark/casedev-cli/commit/e2839ebb80f36deaff2625c4a32e97718ab01126))
* target macOS builds for notarization ([2059ddb](https://github.com/CaseMark/casedev-cli/commit/2059ddb4b126d6eff36e5fc9baee9e15400974ee))
* update casedev-go SDK to v0.6.0 to add Agent V1 Chat service support ([cb7bf0d](https://github.com/CaseMark/casedev-cli/commit/cb7bf0de41e8b3c979688f3df0fe6a88ff44acf9))


### Chores

* add macOS release notarization ([7b74b03](https://github.com/CaseMark/casedev-cli/commit/7b74b035e6c8dae769e3684f47315b513ce9df46))
* **ci:** skip uploading artifacts on stainless-internal branches ([d938a04](https://github.com/CaseMark/casedev-cli/commit/d938a04191c420b733bb99b03b0b5985e535c12a))
* **internal:** codegen related update ([efe2292](https://github.com/CaseMark/casedev-cli/commit/efe229244b0a53368a9b47d2ca0386d85113eab0))
* **internal:** codegen related update ([66a3173](https://github.com/CaseMark/casedev-cli/commit/66a31733024c44f38b0e6ab46e0e4d635c4dd458))
* **internal:** codegen related update ([05a58cd](https://github.com/CaseMark/casedev-cli/commit/05a58cd1393c8798cd1d1bb391894321f3a0c314))
* **internal:** codegen related update ([d5f0216](https://github.com/CaseMark/casedev-cli/commit/d5f02164bdf8d0f0b8b36db44df5dfa5d814ae5f))
* **internal:** codegen related update ([dc4b3bb](https://github.com/CaseMark/casedev-cli/commit/dc4b3bb3fa3c271b252b3fe0219f4368c2f5daf2))
* **internal:** codegen related update ([f29b016](https://github.com/CaseMark/casedev-cli/commit/f29b016b102709b1d699ca891e1d18400a0690d7))
* **internal:** codegen related update ([f991970](https://github.com/CaseMark/casedev-cli/commit/f991970d49115232651c9785d02e590cd366bdda))
* **internal:** codegen related update ([d7dee7d](https://github.com/CaseMark/casedev-cli/commit/d7dee7d7e67a99d072e82b4a2745e13c53e2de7f))
* **internal:** codegen related update ([cb9f807](https://github.com/CaseMark/casedev-cli/commit/cb9f807ae0e253dabefbf7dc659d2354eee6b6f4))
* **internal:** codegen related update ([2e8f723](https://github.com/CaseMark/casedev-cli/commit/2e8f72335f3fc89aa3821fcb826af32eb88bd974))
* **internal:** codegen related update ([a27a6d2](https://github.com/CaseMark/casedev-cli/commit/a27a6d23a0d78a1b09264c26a826ca9015a8f439))
* **internal:** remove mock server code ([6b046c7](https://github.com/CaseMark/casedev-cli/commit/6b046c779afbab408f53509c75d006e01d9bea96))
* **internal:** tweak CI branches ([684f1eb](https://github.com/CaseMark/casedev-cli/commit/684f1eb5404e4849e964b8b96491623bfed72125))
* replace PR open Slack alert with CI failure alert ([b414a29](https://github.com/CaseMark/casedev-cli/commit/b414a29ba634fe75e5fdf2b8c33b4e32f9d55f23))
* **test:** do not count install time for mock server timeout ([04d9879](https://github.com/CaseMark/casedev-cli/commit/04d9879e22b8a8cdcf113b63167ee38415f4a348))
* update mock server docs ([5eefc0f](https://github.com/CaseMark/casedev-cli/commit/5eefc0ff576f9fc46026740219fd20f0198fd0f1))
* update readme with better instructions for installing with homebrew ([90b772a](https://github.com/CaseMark/casedev-cli/commit/90b772a0393b3eba09af6e85334462f831990006))
* update SDK settings ([c780277](https://github.com/CaseMark/casedev-cli/commit/c780277e2e5aebd28af7e123305e06d3c2f0c233))
* update SDK settings ([e79e380](https://github.com/CaseMark/casedev-cli/commit/e79e38056f07c5a67d0f340fef97b829a99a2629))
* update SDK settings ([80009a7](https://github.com/CaseMark/casedev-cli/commit/80009a7b0313730d8f6e2ca366fc66aed0b304b0))
* update SDK settings ([5c433ee](https://github.com/CaseMark/casedev-cli/commit/5c433eef255aa1bdea5b63113a2f3b4bebe72c03))
* update SDK settings ([b0396d5](https://github.com/CaseMark/casedev-cli/commit/b0396d5e5f6d4525e1d8bb0969543616b9b5be92))
* zip READMEs as part of build artifact ([e825863](https://github.com/CaseMark/casedev-cli/commit/e825863966a1910c8761141fe7bfc82ed95ba0e7))

## 0.21.1 (2026-03-19)

Full Changelog: [v0.21.0...v0.21.1](https://github.com/CaseMark/casedev-cli/compare/v0.21.0...v0.21.1)

### Bug Fixes

* improve linking behavior when developing on a branch not in the Go SDK ([5cae562](https://github.com/CaseMark/casedev-cli/commit/5cae5621bd05cdd2ed4501626a54d556e7106842))

## 0.21.0 (2026-03-19)

Full Changelog: [v0.20.2...v0.21.0](https://github.com/CaseMark/casedev-cli/compare/v0.20.2...v0.21.0)

### Features

* **api:** api update ([302e01a](https://github.com/CaseMark/casedev-cli/commit/302e01af3df9fe604bfcde198c03f1beaca17771))


### Bug Fixes

* clear quarantine for homebrew installs ([18af47f](https://github.com/CaseMark/casedev-cli/commit/18af47f8fb8aeefde6aefae39613637f450c65fa))
* target macOS builds for notarization ([2059ddb](https://github.com/CaseMark/casedev-cli/commit/2059ddb4b126d6eff36e5fc9baee9e15400974ee))


### Chores

* **internal:** codegen related update ([efe2292](https://github.com/CaseMark/casedev-cli/commit/efe229244b0a53368a9b47d2ca0386d85113eab0))

## 0.20.2 (2026-03-18)

Full Changelog: [v0.20.1...v0.20.2](https://github.com/CaseMark/casedev-cli/compare/v0.20.1...v0.20.2)

### Chores

* add macOS release notarization ([7b74b03](https://github.com/CaseMark/casedev-cli/commit/7b74b035e6c8dae769e3684f47315b513ce9df46))

## 0.20.1 (2026-03-18)

Full Changelog: [v0.20.0...v0.20.1](https://github.com/CaseMark/casedev-cli/compare/v0.20.0...v0.20.1)

### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([72e551b](https://github.com/CaseMark/casedev-cli/commit/72e551bee1115cc80c456f1aa60090837f064c8f))

## 0.20.0 (2026-03-17)

Full Changelog: [v0.19.0...v0.20.0](https://github.com/CaseMark/casedev-cli/compare/v0.19.0...v0.20.0)

### Features

* **api:** api update ([9abcbc3](https://github.com/CaseMark/casedev-cli/commit/9abcbc3ac32c890b8e800100ec1179a910634421))

## 0.19.0 (2026-03-17)

Full Changelog: [v0.18.1...v0.19.0](https://github.com/CaseMark/casedev-cli/compare/v0.18.1...v0.19.0)

### Features

* **api:** api update ([55398aa](https://github.com/CaseMark/casedev-cli/commit/55398aaee864618bcbbe0823c8e283475b060a5b))


### Bug Fixes

* better support passing client args in any position ([1abc117](https://github.com/CaseMark/casedev-cli/commit/1abc1174fd9e3f4bfd069c139e88d485bd2666fd))
* improved workflow for developing on branches ([561bc6e](https://github.com/CaseMark/casedev-cli/commit/561bc6e0671fb1b79d86a8e59f3847edac60fcf3))
* no longer require an API key when building on production repos ([27607fe](https://github.com/CaseMark/casedev-cli/commit/27607fecc4a04348461621afd8abf95ca74260d5))


### Chores

* **internal:** codegen related update ([66a3173](https://github.com/CaseMark/casedev-cli/commit/66a31733024c44f38b0e6ab46e0e4d635c4dd458))
* **internal:** tweak CI branches ([684f1eb](https://github.com/CaseMark/casedev-cli/commit/684f1eb5404e4849e964b8b96491623bfed72125))
* update SDK settings ([c780277](https://github.com/CaseMark/casedev-cli/commit/c780277e2e5aebd28af7e123305e06d3c2f0c233))

## 0.18.1 (2026-03-14)

Full Changelog: [v0.18.0...v0.18.1](https://github.com/CaseMark/casedev-cli/compare/v0.18.0...v0.18.1)

### Bug Fixes

* only set client options when the corresponding CLI flag or env var is explicitly set ([fd1ab05](https://github.com/CaseMark/casedev-cli/commit/fd1ab052cd8127b05b159aa2a1b0ab931cdd9d62))

## 0.18.0 (2026-03-13)

Full Changelog: [v0.17.1...v0.18.0](https://github.com/CaseMark/casedev-cli/compare/v0.17.1...v0.18.0)

### Features

* **api:** api update ([7d0fdb9](https://github.com/CaseMark/casedev-cli/commit/7d0fdb98e3e83a7c66da614d49e8f04b132490c1))

## 0.17.1 (2026-03-13)

Full Changelog: [v0.17.0...v0.17.1](https://github.com/CaseMark/casedev-cli/compare/v0.17.0...v0.17.1)

### Bug Fixes

* add casedev-go checksum entries ([#46](https://github.com/CaseMark/casedev-cli/issues/46)) ([11c0b42](https://github.com/CaseMark/casedev-cli/commit/11c0b424578eca073b5b73b54df9f4aa50fe74ee))

## 0.17.0 (2026-03-13)

Full Changelog: [v0.16.1...v0.17.0](https://github.com/CaseMark/casedev-cli/compare/v0.16.1...v0.17.0)

### Features

* **api:** api update ([e06c732](https://github.com/CaseMark/casedev-cli/commit/e06c7324a2defa3ff645a8a7d83cf6b30be22b20))

## 0.16.1 (2026-03-13)

Full Changelog: [v0.16.0...v0.16.1](https://github.com/CaseMark/casedev-cli/compare/v0.16.0...v0.16.1)

### Chores

* **internal:** codegen related update ([05a58cd](https://github.com/CaseMark/casedev-cli/commit/05a58cd1393c8798cd1d1bb391894321f3a0c314))

## 0.16.0 (2026-03-12)

Full Changelog: [v0.15.1...v0.16.0](https://github.com/CaseMark/casedev-cli/compare/v0.15.1...v0.16.0)

### Features

* **api:** api update ([f5c04cd](https://github.com/CaseMark/casedev-cli/commit/f5c04cdd99d24af359d313d85c5124b14bde4d8b))

## 0.15.1 (2026-03-12)

Full Changelog: [v0.15.0...v0.15.1](https://github.com/CaseMark/casedev-cli/compare/v0.15.0...v0.15.1)

### Bug Fixes

* fix for test cases with newlines in YAML and better error reporting ([8735c70](https://github.com/CaseMark/casedev-cli/commit/8735c70b6a32e7abf9252c94a14d76f0a70204d3))

## 0.15.0 (2026-03-12)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/CaseMark/casedev-cli/compare/v0.14.0...v0.15.0)

### Features

* **api:** api update ([0673986](https://github.com/CaseMark/casedev-cli/commit/0673986c650e889d56636668b7b7434838b279f4))


### Chores

* **internal:** codegen related update ([d5f0216](https://github.com/CaseMark/casedev-cli/commit/d5f02164bdf8d0f0b8b36db44df5dfa5d814ae5f))

## 0.14.0 (2026-03-11)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/CaseMark/casedev-cli/compare/v0.13.0...v0.14.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([161d0f2](https://github.com/CaseMark/casedev-cli/commit/161d0f2f88a18632f6b01edb2c1a10ad8f5186aa))
* **api:** api update ([bbf3531](https://github.com/CaseMark/casedev-cli/commit/bbf3531a2bddc2d903a3ceb8c1bd9972220826a4))
* **api:** api update ([d530243](https://github.com/CaseMark/casedev-cli/commit/d5302438f10ebf09d244939e0a2f906fb7e75e3e))


### Bug Fixes

* fix for encoding arrays with `any` type items ([66aca04](https://github.com/CaseMark/casedev-cli/commit/66aca04587d2e606a8c90ff5e10f0b64dea4b3b2))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([d938a04](https://github.com/CaseMark/casedev-cli/commit/d938a04191c420b733bb99b03b0b5985e535c12a))
* **internal:** codegen related update ([dc4b3bb](https://github.com/CaseMark/casedev-cli/commit/dc4b3bb3fa3c271b252b3fe0219f4368c2f5daf2))
* **internal:** codegen related update ([f29b016](https://github.com/CaseMark/casedev-cli/commit/f29b016b102709b1d699ca891e1d18400a0690d7))
* **test:** do not count install time for mock server timeout ([04d9879](https://github.com/CaseMark/casedev-cli/commit/04d9879e22b8a8cdcf113b63167ee38415f4a348))

## 0.13.0 (2026-03-07)

Full Changelog: [v0.12.1...v0.13.0](https://github.com/CaseMark/casedev-cli/compare/v0.12.1...v0.13.0)

### Features

* support passing required body params through pipes ([e61b278](https://github.com/CaseMark/casedev-cli/commit/e61b27882931978ef068b50f759b2c22c0fa965d))

## 0.12.1 (2026-03-07)

Full Changelog: [v0.12.0...v0.12.1](https://github.com/CaseMark/casedev-cli/compare/v0.12.0...v0.12.1)

### Chores

* **internal:** codegen related update ([f991970](https://github.com/CaseMark/casedev-cli/commit/f991970d49115232651c9785d02e590cd366bdda))

## 0.12.0 (2026-03-06)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/CaseMark/casedev-cli/compare/v0.11.0...v0.12.0)

### Features

* **api:** api update ([cc9a197](https://github.com/CaseMark/casedev-cli/commit/cc9a197be9209d744372909c690b222232b48949))
* **api:** api update ([57dc9a8](https://github.com/CaseMark/casedev-cli/commit/57dc9a8f52e1a0ffc0219d6d7cc69323d1fd57f7))

## 0.11.0 (2026-03-05)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/CaseMark/casedev-cli/compare/v0.10.0...v0.11.0)

### Features

* **api:** api update ([c3b8f51](https://github.com/CaseMark/casedev-cli/commit/c3b8f5128945bc7eaff19e42659eeea0f84f7e5a))


### Bug Fixes

* bump casedev-go from v0.8.0 to v0.9.0 ([#36](https://github.com/CaseMark/casedev-cli/issues/36)) ([e370249](https://github.com/CaseMark/casedev-cli/commit/e37024979013b8479d16ed24e36fdfad27bbd60f))

## 0.10.0 (2026-03-05)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/CaseMark/casedev-cli/compare/v0.9.0...v0.10.0)

### Features

* **api:** api update ([a99754f](https://github.com/CaseMark/casedev-cli/commit/a99754fd54fe7104c2eebba99af7b70318c7adab))

## 0.9.0 (2026-03-04)

Full Changelog: [v0.8.1...v0.9.0](https://github.com/CaseMark/casedev-cli/compare/v0.8.1...v0.9.0)

### Features

* **api:** api update ([fc663ad](https://github.com/CaseMark/casedev-cli/commit/fc663adddf3eca77a8e658de11de42eb4952c79d))

## 0.8.1 (2026-03-04)

Full Changelog: [v0.8.0...v0.8.1](https://github.com/CaseMark/casedev-cli/compare/v0.8.0...v0.8.1)

### Bug Fixes

* avoid printing usage errors twice ([1f49bd9](https://github.com/CaseMark/casedev-cli/commit/1f49bd9c3a000f057be6ba88b400b2bd11dea74a))

## 0.8.0 (2026-03-04)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/CaseMark/casedev-cli/compare/v0.7.0...v0.8.0)

### Features

* add support for file downloads from binary response endpoints ([ca496e8](https://github.com/CaseMark/casedev-cli/commit/ca496e865692771ab3dfb1f99f97382e25ac3fc6))

## 0.7.0 (2026-03-04)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/CaseMark/casedev-cli/compare/v0.6.0...v0.7.0)

### Features

* improved documentation and flags for client options ([94e1cb9](https://github.com/CaseMark/casedev-cli/commit/94e1cb9ba8b19d54e2729ed89b9869659ff1e62f))


### Chores

* **internal:** codegen related update ([d7dee7d](https://github.com/CaseMark/casedev-cli/commit/d7dee7d7e67a99d072e82b4a2745e13c53e2de7f))
* **internal:** codegen related update ([cb9f807](https://github.com/CaseMark/casedev-cli/commit/cb9f807ae0e253dabefbf7dc659d2354eee6b6f4))

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
