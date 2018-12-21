??? tip "naming"
    Check for consistent naming conventions within components of the spec.

    The supported naming conventions are:

    - PascalCase
    - CamelCase
    - KebabCase
    - SnakeCase

    ```yaml
    rules:
      naming:
        paths: KebabCase
        tags: PascalCase
        operation: CamelCase
        parameters: SnakeCase
        definitions: PascalCase
        properties: camelCase
    ```

??? tip "noEmptyDescriptions"
    Disallows empty descriptions for various components of the spec.

    A value of `true` will cause empty or missing descriptions to become errors.
    `false` or omitted will allow empty descriptions.

    ```yaml
    rules:
      noEmptyDescriptions:
        operations: true
        properties: true
        parameters: false
    ```


??? tip "noEmptyOperationIDs"
    Disallows empty operation IDs

    !!! example

        ```yaml
        rules:
          noEmptyOperationIDs: true
        ```

    !!! success "Good"

        ```yaml
        paths:
          /pets:
            get:
              operationId: listPets
              ...
        ```

    !!! error "Bad"

        ```yaml
        paths:
          /pets:
            get:
              operationId: ""
              ...
        ```


??? tip "slashTerminatedPaths"

    Check that paths consistently end with a slash or not.

    A value of `true` requires paths to end in a slash.
    A value of `false` will require all paths to **not** end with a slash.

    !!! example

        ```yaml
        rules:
          slashTerminatedPaths: true
        ```

    !!! success "Good"

        ```yaml
        paths:
          /pets/:
            ...
        ```

    !!! error "Bad"

        ```yaml
        paths:
          /pets:
            ...
        ```

??? tip "noEmptyTags"

    Check that all operations have at least 1 non-empty tag

    !!! example

        ```yaml
        rules:
          noEmptyTags: true
        ```

    !!! success "Good"

        ```yaml
        paths:
          /pets/:
            get:
              tags:
                - pets
              ...
        ```

    !!! error "Bad"

        ```yaml
        paths:
          /pets/:
            get:
              tags: []
              ...
        ```

??? tip "noUnusedDefinitions"

    Disallows definitions that aren't used anywhere

    !!! example

        ```yaml
        rules:
          noUnusedDefinitions: true
        ```

    !!! success "Good"

        ```yaml
        paths:
          /pets/:
            responses:
              200:
                schema:
                  $ref: '#/definitions/User'

        definitions:
          User:
            type: object
              ...
        ```

    !!! error "Bad"

        ```yaml
        paths:
          /pets/:
            responses:
              200:
                schema:
                  type: object

        definitions:
          User:
            type: object
        ```

??? tip "noDuplicateOperationIDs"

    Disallows duplicate operation IDs.

    !!! example

        ```yaml
        rules:
          noDuplicateOperationIDs: true
        ```

    !!! success "Good"

        ```yaml
        paths:
          /pets/:
            operationId: listPets
          /pets/{id}/:
            operationId: getPet
        ```

    !!! error "Bad"

        ```yaml
        paths:
          /pets/:
            operationId: listPets
          /pets/{id}/:
            operationId: listPets
        ```
