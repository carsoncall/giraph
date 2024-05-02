# Giraph
### A codebase comprehension tool that leverages graphs to supercharge your work

Feature Goals:
1. Show relationships between code files (using import statements). This is the simplest, and the thing that has been mostly implmented thus far.
2. Show inheritance relationships. This functionality already exists in some form in other tools.
3. Show function stacks.
4. Show data flow (related to 3)
5. Eventually, add a configuration interface and a linter so that project managers can define engineering rules/best practices etc and this tool can enforce the rules (like a "design linter")

Components:
- Database
  - right now, it uses Neo4J which comes with convenient graph visualization tools
  - In the future, using an embedded, lighter-weight database would be better
- Parser
  - written in Go
  - handles using the tree-sitter parser to do things with your codebase
  - in the future, the database needs to be abstracted using a DAO or interface etc
  - Adding support for more languages
- Frontend:
  - Right now, the neo4j-browser serves as a frontend
  - We need an interface to allow for different frontends, from React Flow to cli (bubble-tea?) to VSCode
