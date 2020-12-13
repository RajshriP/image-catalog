# Image Catalog

## How to run?
`go run main.go`

## APIs
The APIs will always return JSON response. In case of errors, it will have the following structure:
```json
{
  "error": "actual error message"
}
```

### Upload Image API
URL: `/api/v1/upload/`\
Method: POST (multipart form)

Body:\
`image`: image file

Response:
```json
{
  "id": 123,
  "path": "/images/abc.jpg"
}
```


### Get Images API
This API will return the list of images sorted by the upload time, in descending order.

URL: `/api/v1/images/`\
Method: GET

Params:\
`page_no`: Page number (optional, will default to 1 in case of invalid values)\
`per_page`: Number of images per page (optional, will default to 10 in case of invalid values)

Response:
```json
{
  "page_no": 1,
  "per_page": 10,
  "pages": 1,
  "data": [
    {"id": 3, "path": "/images/c.jpg"},
    {"id": 2, "path": "/images/b.jpg"},
    {"id": 1, "path": "/images/a.jpg"}
  ]
}
```


### Get Image API
This API will return the image path for the given ID

URL: `/api/v1/image/`\
Method: GET

Params:\
`id`: Image ID

Response:
```json
{
  "id": 123,
  "path": "/images/abc.jpg"
}
```