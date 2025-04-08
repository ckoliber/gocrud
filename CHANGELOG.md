## [1.2.1](https://github.com/ckoliber/gocrud/compare/v1.2.0...v1.2.1) (2025-04-08)


### Bug Fixes

* **deps:** update module github.com/ckoliber/gocrud to v1.2.0 ([86d6252](https://github.com/ckoliber/gocrud/commit/86d625247e51675d8f36710d3dfc0071ce315dc2))

# [1.2.0](https://github.com/ckoliber/gocrud/compare/v1.1.0...v1.2.0) (2025-04-08)


### Features

* add docs pages ([b46e728](https://github.com/ckoliber/gocrud/commit/b46e728ab51af8a66273dcc6d12787caf45cad77))

# [1.1.0](https://github.com/ckoliber/gocrud/compare/v1.0.0...v1.1.0) (2025-04-08)


### Features

* add docs ([43761bb](https://github.com/ckoliber/gocrud/commit/43761bb08516b0f15e26a4abd37003030979a1e0))

# 1.0.0 (2025-04-08)


### Bug Fixes

* change CI token ([aa9b742](https://github.com/ckoliber/gocrud/commit/aa9b742b2bc059bbb2d973115da8e50ff8ae86b1))
* change Post method ([9539b9b](https://github.com/ckoliber/gocrud/commit/9539b9b93697c0682be29f9af1caf4bd74864de1))
* change tag format ([db93fa0](https://github.com/ckoliber/gocrud/commit/db93fa0fe18f13f85664239b9175e18a00335d15))
* panic problem in WhereToString ([b26e57c](https://github.com/ckoliber/gocrud/commit/b26e57cccf15bf17fb714485d0d2d95727e50a66))
* remove unused dependencies ([d9a758d](https://github.com/ckoliber/gocrud/commit/d9a758de88554716022c13b91d6bd3e640395fa1))
* skip key fields in Set method ([7fed676](https://github.com/ckoliber/gocrud/commit/7fed676375869cb4bdc5b14b590002a8149f400c))


### Features

* add CI ([33d8593](https://github.com/ckoliber/gocrud/commit/33d85938269dc1128bd38e1caac9ac84132f240b))
* add controller, repository, schema internal modules ([e0740a7](https://github.com/ckoliber/gocrud/commit/e0740a75088c15c51d611e3aaee64659b09dab83))
* add dynamic Register params from model spec ([0e2cda2](https://github.com/ckoliber/gocrud/commit/0e2cda2376e08caf1846882b16f08eccfe0d8bf7))
* add gocrud Register method ([2800466](https://github.com/ckoliber/gocrud/commit/2800466de243b5e189138c110a77c077460d4e4f))
* add internal packages ([2149c30](https://github.com/ckoliber/gocrud/commit/2149c302a1257722961d0c4276cada1cd8d0b361))
* add logs, comments ([3d0e823](https://github.com/ckoliber/gocrud/commit/3d0e8232cdb68ac0b83980f403082086f27401b1))
* add new api internal module ([af45d1f](https://github.com/ckoliber/gocrud/commit/af45d1f1cb465fc445db42f0f9a16c3b2cc25c00))
* add order, where query deepObjects ([86de5fe](https://github.com/ckoliber/gocrud/commit/86de5fe4763be8beb1b78a8b6018b947aca191ca))
* add PostSelect method ([e146d02](https://github.com/ckoliber/gocrud/commit/e146d02823108fc307c4093b3036021bc530296a))
* add put_single id parameter ([efcd900](https://github.com/ckoliber/gocrud/commit/efcd9003461820e41a85dcb301a2b1624064820e))
* add request context to repository and hooks ([98e3261](https://github.com/ckoliber/gocrud/commit/98e3261c239ddabb397a612f025b87218075f4d7))
* add SQL templates to CRUDRepository ([efcb735](https://github.com/ckoliber/gocrud/commit/efcb735f637631e380415fa9963f458f912bb058))
* add support for array of string type condition operator values ([f01ad91](https://github.com/ckoliber/gocrud/commit/f01ad91ccb554811499c264a2e975bb681e4d2a3))
* add transactional put ([65f628b](https://github.com/ckoliber/gocrud/commit/65f628bc01f60f190f3a40d994014166c9c4fe23))
* add where schema ([2a24a4f](https://github.com/ckoliber/gocrud/commit/2a24a4f460d69f3a8d9704f39ee8eb5640ff473e))
* change folder structure ([414e685](https://github.com/ckoliber/gocrud/commit/414e68541c754c55dc6e3c9c04fcee3cc4e8fd4d))
* change internal modules structure ([ad01d1d](https://github.com/ckoliber/gocrud/commit/ad01d1d6ae12bf9dd0573ab0b8613445fb3e9129))
* fix MSSQL syntax problems ([cf1c919](https://github.com/ckoliber/gocrud/commit/cf1c9198c8e63d124c24e4849757b24e7b0211cc))
* implement cross db CRUD queries for delete, put, post ([9f95d27](https://github.com/ckoliber/gocrud/commit/9f95d27c7b0f6eab10c720fd17ac1c439f33e085))
* implement CRUDRepository methods ([3f65ab2](https://github.com/ckoliber/gocrud/commit/3f65ab20b8140e8ede071f16f11e7706b6e69377))
* implement CRUDRepository using go-sqlbuilder ([634b6c7](https://github.com/ckoliber/gocrud/commit/634b6c7af7b1d30db93276c71bd71031fc6c6296))
* implement CRUDService ([8a0ba22](https://github.com/ckoliber/gocrud/commit/8a0ba22de64bc3a97b8c40bdcb9d3d10098b1c01))
* implement Fields, Values, Set methods of SQLBuilder ([74aca34](https://github.com/ckoliber/gocrud/commit/74aca348b1e3ed7ad387d75460ab6d76c1e44929))
* implement generic SQLBuilder ([174e7a0](https://github.com/ckoliber/gocrud/commit/174e7a037f9ef7a2c4947b357845e91cb263ad86))
* implement MySQL driver ([3133c6d](https://github.com/ckoliber/gocrud/commit/3133c6d3cafbedae063e19c06c7466c187c70ccc))
* implement order, where schema validations ([2cefb69](https://github.com/ckoliber/gocrud/commit/2cefb694158380260d7bc39529f4d4456f1b87be))
* implement postgres dialect Put method ([c9b7e8f](https://github.com/ckoliber/gocrud/commit/c9b7e8fa3dd5ad552f1e4f8996269bad9ac784b7))
* implement SQLite, MSSQL drivers ([300045e](https://github.com/ckoliber/gocrud/commit/300045ef6abe04dcfaaa2c8ea0837de9ba5e848d))
* init project ([5caf823](https://github.com/ckoliber/gocrud/commit/5caf823664d3d8b3f60eaf7f13f695be834de199))
* **mysql:** implement Post using LastInsertId and RowsAffected ([3056c65](https://github.com/ckoliber/gocrud/commit/3056c6581ddbefe1411ec4f5fcd5119be779843c))
* **mysql:** implement two stage Put, Delete methods ([7d5ed9d](https://github.com/ckoliber/gocrud/commit/7d5ed9dcefa953e2080c73bbb91cd2907a4c1d64))
* remove PATCH method (use autopatch) ([9c3dcdb](https://github.com/ckoliber/gocrud/commit/9c3dcdb6a5cf7c9316f7771d27569187cc40c22d))
