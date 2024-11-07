# Goland and Gin

This is patlytics's backend of interview
## Technologies Used

- Golang 1.23.2
- Gin

## How to Use

You need to apply a Claude AI's Api key and input on it
```bash
echo API_KEY="{key}" > .env
```

```bash
go mod tidy
```

### Run the development server

```bash
go run main.go
```


1. /search?keyword={keyword}&&type={type}
   You can fast search information by company name or patent id ,and This api support fuzzy query
    type's value accepts "company" and "patent" only

2. /report?company={company name}&&patent={patent id }
   You can use it to analyze by company name and patent id 

### You also can use docker to test this project

