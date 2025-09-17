# TODO

## Core Features

- [x] **Query**  
       _Ask a natural language question and get an answer._  
       **Example:**  
       `umm how to list files`

- [x] **Run Last Command**  
       _Re-run the most recent suggested command._  
       **Example:**  
       `umm --run`

- [x] **Follow-Up Command**  
       _Ask a follow-up question using the most recent query as context._  
       **Example:**  
       `umm + "what about with curl?"`

- [ ] **Contextual History Search**  
       _Search through previous interactions by keywords._  
       **Example:**  
       `umm history --search "curl"`

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
