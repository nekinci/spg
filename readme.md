# spg 

![ci](https://github.com/nekinci/spring-profile-generator/actions/workflows/ci.yaml/badge.svg)

`spg` is a tool that generates Spring profiles based on the current environment. 

If you develop your micro services with real environments and your yaml files not up to date probably you spend time fixing.
The project aims to save you time and you can use it during development and testing.

Example:
You have a `application.yaml` file with the following content:
    
    server:
        port: 8080
    client:
        urls:
          serviceA: servicea.test.com
          serviceB: serviceb.test.com

Let's assume you want debug the application in prp environment. But you don't have a `application-prp.yaml` file. And you don't know prp urls. 
Probably you don't want to create a new file. Because your job is to debug or develop your application. With `spg` you can generate a `application-prp.yaml` file with the following content:
    
    server:
        port: 8080
    client:
        urls:
          serviceA: servicea.prp.com
          serviceB: serviceb.prp.com


But first you need to define trainer config file that contains url information of your services. Because spg don't know the url of your services. The following config file is an example:

    version: v1
    information:
        absolute-configs:
          - config-key: spring.http.encoding.force # means change the value of spring.http.encoding.force to true
            environment:
                oc: true
                test: false
                prp: true
                prod: false
          - config-key: keycloak.securityConstraints[0].authRoles[0] # means change the first role of the first security constraint
            environment:
                oc: user
                test: user
                prp: admin
                prod: admin
          - config-key: keycloak.roles[].value # means change all role values under keycloak.roles
            environment:
                oc: user
                test: user
                prp: admin
                prod: admin
        fields:
            - keys: # metadata keys for decide which fields to change
                - a-service
                - aService
              type: url
              environment:
                oc: 
                  scheme: http
                  value: a-service:8080
                test:
                  scheme: http
                  value: a-service-test.cloud.com
                prp:
                  scheme: http
                  value: a-service-prp.cloud.com
                prod:
                  scheme: https
                  value: a-service.com


#### TODO

* [x] Init project
* [x] Implement Pretty Print Map
* [x] Implement Merge Map with override strategy
* [x] Declare Trainer Yaml Struct
* [x] Add CLI Commands
* [x] Implement generate method
* [ ] Add wildcard feature for absolute configs
* [ ] Add Usage Manuel
* [x] Add Tests
* [ ] Add Go Releaser
* [ ] Publish with Homebrew, etc.
* [x] Add Github Actions
* [x] Add Dockerfile


### Known Issues

- **Don't handled yet metadata in trainer yaml file** (e.g. `information.fields.keys`)
- **Don't handled yet wildcard feature** (e.g. `absolute-configs.config-key`)
- **Config array selectors only support for primitive types. Don't supported for complex types.** (e.g. `maps or arrays`) 