## TODO

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

- mal-cli anime [ranking|seasonal] "search title"
    - user will see the list of anime matching the search (fzf)
    - on selecting one of them, (if options are given) they will get the details of the selection


## Short Term Goal

complete the public routes handling and then create a client capable of requesting data from those public routes

#### Script Commands
1. search

> [!NOTE]
> Libraries for creating the terminal client
> - cobra for all the functionality.
> - tview for user interaction and data display. 
> - fzf (probably) for choosing from a list of data.




## Today's Task

design how the user will use the script...
- default script behaviour would be to search for anime.
    - searching for an anime will give them with a fzf list. 
- to be explicit about searching "mal-cli search "name" "
