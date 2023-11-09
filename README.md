# stockels

Personalize stocks analytic data

Key Feature:

- Stock information with parameter that can adjust by your self
- Whitelist stock to get informed whenever you open the app
- Generate report file to recap all your analytic data
- Get technical analytic for each stock to determine support and resistance price per quarter
- Get current fundamental analytic data for every stock on IDX
- Latest business news to support your analysis

* Get improvement of company profit per quarter

Technical Feature:

- Build in golang with Gin framework and GORM ORM
- Use go routines to fetch data concurrently from third party API
- Using Redis to cache data from goapi.id and serpapi.com which only has limit 100 req/month
- Using S3 bucket to store report data in .csv file format
- Schema first GraphQL api for improve dev experience and minimalize data transfer between frontend and backend
- Unit and Integration test for all importent service and endpoint to make sure everything works as it should be
