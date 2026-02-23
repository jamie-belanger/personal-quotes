# Roadmap
Here are my current plans for this project

## Phase 1: MVP Groundwork
An application that allows the user to maintain and use a personal library of quotes.

<details>
<summary>Phase 1 Tasks = In Progress</summary>

- ✅ **Basic Structure**
  - ✅ Um, everything I did before making this roadmap (whew!)
  - ✅ HTML sanitizing
- ♻️ **Database schema design**
  - ✅ Quote = for storing quotes
  - 🔲 Tag = for storing tags, as metadata for quotes (ie "funny" or "motivational")
  - 🔲 QuoteTags = bridge table for linking Quotes to Tags
  - 🔲 QuoteUser = for storing user information
- ✅ **REST API**
  - ✅ Controllers for Quotes
  - 🔲 Controllers for Tags
  - 🔲 Quotes search endpoint &mdash; GET all, author search, body search, paging logic, filter by tag, etc
  - 🔲 Swagger for documentation?
  - ✅ Test API with curl
- ✅ **SQLite database driver**

</details>


## Phase 2: Full MVP
User and developer experience, resulting in a fully usable program.

<details>
<summary>Phase 2 Tasks</summary>

- 🔲 **Unit Tests** for server-side library code
- 🔲 **Database**
  - 🔲 MySQL database driver
  - 🔲 Better way to handle database seed data
  - 🔲 Better way to handle similar boilerplate by database?
  - 🔲 Postgres database driver
  - 🔲 User account table
  - 🔲 Add "Last Update" logic
- 🔲 **Authentication**
  - 🔲 User accounts controller
  - 🔲 Initial admin user created as part of initial database seed
  - 🔲 Basic user/pass auth
  - 🔲 Signups
  - 🔲 Optional OIDC auth
  - 🔲 JWTs for identity tokens
  - 🔲 Secure API endpoints with auth checks (except anon endpoint)
- 🔲 **Application GUI** = Svelte? or something else?
  - 🔲 User account / preferences page
  - 🔲 Time zone handling
  - 🔲 Test login process and preferences with multiple users
  - 🔲 Tag setup screen
    - 🔲 List view
    - 🔲 Full view with CRUD
  - 🔲 Quotes
    - 🔲 List view
    - 🔲 Full view with CRUD
  - 🔲 Page states (404, 403, etc)
  - 🔲 Dashboard
    - 🔲 Database / user stats
    - 🔲 Quote widget for testing
- 🔲 **Docker**
  - 🔲 Docker build
  - 🔲 Docker compose
  - 🔲 Test container on personal server
  - 🔲 Github actions CI/CD pipeline --> GHCR? Docker hub?
- 🔲 Announce on r/selfhosted ?

</details>



## Phase 3: Polish and Play
Add some polish to the application, and start playing with additional technologies.

<details>
<summary>Phase 3 Tasks</summary>

- 🔲 **Documentation**
  - 🔲 Comprehensive API usage examples (CURL, PHP, Typescript, etc)
- 🔲 **Logging** = I have logging code in there, but want to spend more effort here:
  - ✅ JSON logging
  - 🔲 Middleware to enrich based on user context
  - 🔲 Integration with log viewer (Seq? something else?)
- 🔲 Performance testing
- 🔲 **[GraphQL](https://graphql.org/)** endpoint for easier external querying
- 🔲 **Additional Ideas**
  - 🔲 GUI update to find new quotes from author
  - 🔲 Cleanup tools (find duplicate quotes, consolidate tags, etc)

</details>

