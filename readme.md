# Spring Boot Profile Yaml Generator

-------------------------------------

![ci](https://github.com/nekinci/spring-profile-generator/actions/workflows/ci.yaml/badge.svg)


Spring Boot Yaml Generator changes environment urls by given environment.
If you develop your micro services with real environments and your yaml files not up to date probably you spend time fixing.

The project working with trainer yaml file that declares environmental informations. You should just give to existing yaml and specify the environment for generate a new.

#### TODO

* [x] Init project
* [x] Implement Pretty Print Map
* [x] Implement Merge Map with override strategy
* [x] Declare Trainer Yaml Struct
* [ ] Add CLI Commands
* [x] Implement generate method
* [ ] Add wildcard feature for absolute configs
* [ ] Add Usage Manuel
* [x] Add Tests
* [ ] Add Go Releaser
* [ ] Publish with Homebrew, etc.
* [x] Add Github Actions
* [x] Add Dockerfile


### Known Issues

-------

- **Don't handled yet metadata in trainer yaml file** (e.g. `information.fields.keys`)