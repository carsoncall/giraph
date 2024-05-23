# Giraph
### A codebase comprehension tool that leverages graphs to supercharge your work

### Get started
First, fetch the dependencies:

`go mod tidy`

Then, run it:

`go run cmd/main.go path/to/your/codebase`


### Structure hints:
- /cmd: entry point for the application. Code in this directory should be minimized if possible.
- /internal: code that shouldn't be imported (application specific code)
- /pkg: code that could be imported into other applications or libraries (generic code)
- /vendor: where application dependencies are stored
These standards were pulled from the (Golang Standards repo)[https://github.com/golang-standards/project-layout]

Feature Goals:
1. Show relationships between code files (using import statements). This is the simplest, and the thing that has been mostly implmented thus far.
2. Show inheritance relationships. This functionality already exists in some form in other tools.
3. Eventually, add a configuration interface and a linter so that project managers can define engineering rules/best practices etc and this tool can enforce the rules (like a "design linter")
4. Show function stacks.
5. Show data flow (related to 3)


Components:
- Database
  - right now, it uses Neo4J which comes with convenient graph visualization tools
  - In the future, using an embedded, lighter-weight database would be better
  - Or, even better, we could ditch the database entirely and use an in-memory structure with support for serialization/deserialization for when we need to store the data. 
- Parser
  - written in Go
  - handles using the tree-sitter parser to do things with your codebase
  - in the future, the database needs to be abstracted using a DAO or interface etc
  - Adding support for more languages
- Frontend:
  - Right now, the neo4j-browser serves as a frontend
  - We need an interface to allow for different frontends, from React Flow to cli (bubble-tea?) to VSCode
- Server:
  - This will use an API that passes Protobuf objects to maintain type safety. The goal is to keep the frontend and backend separated, so that the frontend can be substituted out. 
