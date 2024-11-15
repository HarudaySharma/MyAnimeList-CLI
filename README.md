## todo

#### all the public routes for MAL_CLI;

- [x] GET anime list
- [x] GET anime details
- [x] GET anime ranking
- [x] GET seasonal anime

#### all the user routes for MAL_CLI;

- [ ] GET suggested anime
- [ ] GET user anime list
- [ ] UPDATE user anime list status
- [ ] DELETE user anime list item

#### User Script (client)

- [ ] Script Commands
    - [x] search
    - [x] seasonal
    - [x] ranking
    - [ ] user specific commands
- [ ] UI
    - [x] mal-cli anime [ranking|seasonal] "search title"
        - [x] user will see the list of anime matching the search (fzf)
        - [x] on selecting one of them, (if options are given) they will get the details of the selection
    - [ ] show anime details on list selection step.


#### Short Term Goal

- [x] complete the public routes handling and then create a client capable of requesting data from those public routes

---

## Script Commands
- search
- seasonal
- ranking


> [!NOTE]
> Libraries for creating the terminal client
> - cobra for all the functionality.
> - tview for user interaction and data display.
> - fzf (probably) for choosing from a list of data.


``` bash
kitten icat --place 28x28@152x3 --clear --align=center https://bleach-anime.com/assets/img/top/main_09.jpg

```
