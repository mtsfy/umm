# TODO

## Core Features

- [x] **Query**  
       _Ask a natural language question and get an answer._  
       **Example:**  
       `umm how to list files`

- [x] **Run Suggested Command**  
       _Run the most recent suggested command._  
       **Example:**  
       `umm --run` _(last suggested)_
      `umm --run 2`

- [x] **Follow-Up Command**  
       _Ask a follow-up question using the most recent query as context._  
       **Example:**  
       `umm + "what about with curl?"`

- [x] **History**  
       _Manage your past interactions._

  - [x] **View**  
         _Display your interactions with optional pagination._  
         **Example:** `umm history`  
         **Example:** `umm history --page 2 --size 10`

  - [x] **Contextual History Search**  
         _Filter previous interactions by keywords._  
         **Example:** `umm history --search "curl"`

  - [x] **Delete**  
         _Remove interactions from history – either all or specific ones._  
         **Example:** `umm history --delete -1` _(-1 for all history)_
        **Example:** `umm history --delete 2`

- [ ] **Command Preview**  
       _Preview suggested commands before execution with extra details._  
       **Example:**  
       `umm preview "list files in current directory"`

- [ ] **Interactive Mode**  
       _Start an interactive shell that maintains session context._  
       **Example:**  
       `umm interactive`
      `umm> [Your initial query]`
      `umm> [Next query in same session]`

## Optional

- [ ] **Alias Support**  
       _Create and manage shortcuts for frequently used commands._  
       **Example:**  
       `umm alias add ls="ls -la"`
      `umm ls`
