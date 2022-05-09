1. Why do we need URL shortening?#
- less space
- Optimize links across devices
- track individual links to analyze audience
- measure ad campaigns’ performance, or hide affiliated original URLs

2. Functional Requirements
- shorter url with unique aliases
- service should redirect to original link
- user should have choice to choose custom url
- link will expire after some time

3. Non functional Requirements
- system should be highly available
- url redirection should be with minimum latency
- links should not be guessable or predictable
- analysis is always helpful
- link should be internally accessible to other services

4. Estimates
- Capacity
- Lets assume a 100:1 read and write ratio for redirection requests and 500M new URL shortenings per month 
- No. of redirections = 100 * 500M = 50B
- Queries per second = 500M / 30*24*3600 ~= 200 URLs/s
- URLs redirection per second = 100*200 URLs/s = 20K/s

- Storage
- Lets assume this for 5 years, 500M for every month,
- total objects we expect to store = 500 * 5 * 12 months = 30 Billion
- assume each objects required 500 bytes
- total storage = 30 * 500 bytes = 15 TB storage

- Bandwidth
- For write requests, we expect 200 URLs every second
- total incoming data = 200 * 500bytes = 100Kb/s
- For read requests, we expect 20K URLs redirections
- total outgoing data = 20K * 500bytes = 10MB/s

- Memory 
- If we want to cache some of hot URLs that are frequently accessed, 80-20 in ratio
- Since we have 20K requests per second, 1.7 billion request per day
- 20K * 3600 * 24 hr = 1.7 billion
- 20 % of 1.7 billion * 500 bytes = 170 GB duplicates will reduce data

5. System APIs
- Create URL : should accept original url and return tiny url
- Delete URL : delete tiny url

6. Database Design
- we need to store billion of records
- each object is less than 1K
- no relationships in objects
- read heavy

- Schema
- URL: originalUrl, creationData, expirationData, userId
- User: name, email, creationDate, lastLogin

- Database which we can use?
- a NoSql store like DynamoDb, Cassandra or MongoDB is better choice, no relationship and Nosql is easier to scale.

7. Algorithms
- https://tinyurl.com/<eight digit unique code>
- 


