# AssertTiDB


## Concepts

+ Data Generator
  + Plain Data Generator: generates plain data (e.g. integer/varchar(256)/bool)
  + SQL Generator: generates sql according to input plain data
+ Executor: Executes input sql and responds with some feedback
+ Assertor: Judges query result according to input sql
+ Runner: Runs executor and judge

## Usage

1. config data generator (e.g. custom one)
2. config executor (e.g. custom one)
3. config assertor (e.g. custom one)
4. config runner (such as number of runnings, parallelism)


## Todo
- [ ] Generate DDL statements
- [ ] Executor logic: refactor the code currently in main function
- [ ] Assertion logic: judge whether a command runs correctly
- [ ] Runner logic: all of it