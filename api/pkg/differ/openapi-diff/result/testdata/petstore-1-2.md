#### What's New

##### `GET` /pet/{petId}

> Find pet by ID

#### What's Deleted

##### `POST` /pet/{petId}

> Updates a pet in the store with form data

#### What's Deprecated

##### `GET` /user/logout

> Logs out current logged in user session

#### What's Modified

##### `DELETE` /pet/{petId}

> Deletes a pet

###### Parameters:

Added: `newHeaderParam` in `header`


##### `GET` /user/login

> Logs user into the system

###### Parameters:

Deleted: `password` in `query`
> The password for login in clear text

###### Response:

Modified response: **200 OK**
> successful operation

* Added header: `X-Rate-Limit-New`

* Deleted header: `X-Rate-Limit`

* Modified header: `X-Expires-After`

##### `GET` /user/logout

> Logs out current logged in user session


##### `POST` /pet/{petId}/uploadImage

> uploads an image for pet

###### Parameters:

Modified: `petId` in `path`
> ID of pet to update, default false


##### `POST` /user

> Create user

###### Request:

Modified content type: `application/json`

* Added property `newUserFeild` (integer)
    > a new user feild demo

* Deleted property `phone` (string)


##### `POST` /user/createWithArray

> Creates list of users with given input array

###### Request:

Modified content type: `application/json`

* Modified items (array):

    * Added property `newUserFeild` (integer)
        > a new user feild demo

    * Deleted property `phone` (string)


##### `POST` /user/createWithList

> Creates list of users with given input array

###### Request:

Modified content type: `application/json`

* Modified items (array):

    * Added property `newUserFeild` (integer)
        > a new user feild demo

    * Deleted property `phone` (string)


##### `GET` /user/{username}

> Get user by user name

###### Response:

Modified response: **200 OK**
> successful operation

* Modified content type: `application/json`
    * Added property `newUserFeild` (integer)
        > a new user feild demo

    * Deleted property `phone` (string)


* Modified content type: `application/xml`
    * Added property `newUserFeild` (integer)
        > a new user feild demo

    * Deleted property `phone` (string)


##### `PUT` /user/{username}

> Updated user

###### Request:

Modified content type: `application/json`

* Added property `newUserFeild` (integer)
    > a new user feild demo

* Deleted property `phone` (string)


##### `PUT` /pet

> Update an existing pet

###### Request:

Deleted content type: `application/xml`

Modified content type: `application/json`

* Added property `newField` (string)
    > a field demo

* Modified property `category` (object)

    * Added property `newCatFeild` (string)

    * Deleted property `name` (string)


##### `POST` /pet

> Add a new pet to the store

###### Parameters:

Added: `tags` in `query`
> add new query param demo

###### Request:

Modified content type: `application/xml`

* Added property `newField` (string)
    > a field demo

* Modified property `category` (object)

    * Added property `newCatFeild` (string)

    * Deleted property `name` (string)

Modified content type: `application/json`

* Added property `newField` (string)
    > a field demo

* Modified property `category` (object)

    * Added property `newCatFeild` (string)

    * Deleted property `name` (string)


##### `GET` /pet/findByStatus

> Finds Pets by status

###### Parameters:

Modified: `status` in `query`
> Status values that need to be considered for filter

###### Response:

Modified response: **200 OK**
> successful operation

* Modified content type: `application/xml`
    * Modified items (array):

        * Added property `newField` (string)
            > a field demo

        * Modified property `category` (object)


* Modified content type: `application/json`
    * Modified items (array):

        * Added property `newField` (string)
            > a field demo

        * Modified property `category` (object)


##### `GET` /pet/findByTags

> Finds Pets by tags

###### Response:

Modified response: **200 OK**
> successful operation

* Modified content type: `application/xml`
    * Modified items (array):

        * Added property `newField` (string)
            > a field demo

        * Modified property `category` (object)


* Modified content type: `application/json`
    * Modified items (array):

        * Added property `newField` (string)
            > a field demo

        * Modified property `category` (object)


