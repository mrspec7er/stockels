# stockels
Personalize stocks analytic data

Key Feature:
- Stock information with parameter that can adjust by your self
- Whitelist stock to get informed whenever you open the app
- Generate report file to recap all your analytic data

Infrastructure:
- Build in golang with Gin framework and GORM ORM
- Using Redis to cache data from goapi.id which only has limit 1000 req/day
- Using S3 bucket to store report data in .csv file format
