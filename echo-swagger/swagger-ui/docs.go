
package swaggerui
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:8080/swagger-ui",
    "apis": [
        {
            "path": "/hello",
            "description": "hello"
        }
    ],
    "info": {
        "title": "echo-swagger",
        "description": "echo-swagger sample api"
    }
}`
var apiDescriptionsJson = map[string]string{"hello":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:8080",
    "resourcePath": "/hello",
    "apis": [
        {
            "path": "/hello",
            "description": "Hello 'your name'",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Hello",
                    "type": "string",
                    "items": {},
                    "summary": "Hello 'your name'",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "name",
                            "description": "名前",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Hello World",
                            "responseType": "object",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ]
}`,}
