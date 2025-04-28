# moneybag

`moneybag` is a simple and minimal app based on transferring between so-called
bags.

## Some theory

A **bag** is a named unsigned integer variable. And that's all about it, it
just has a name and a balance that is a natural (better say non-negative since
it can be also zero) integer number. Since it's a *money*bag the thing the bag
contains is called **money**.
- There is also a special **"bottomless"** bag that is always available. You
can take infinite amounts of money and put it there too.

A **tag** is a named list of bag names so you can categorize your bags by
tagging them and then addressing to them with this tag.\
A **transfer** is an action of taking a particular **amount** of money from
one bag (**source**) to another (**destination**). If source does not exist,
its balance is considered to be zero. If destination does not exist, it gets
created. If the given amount is more than balance of the source then the
transfer does not make any changes.\
A **book** is an ordered list of transfers. The whole state of all your bags
can be reproduced by re-applying all of the transfers from scratch. The book
contains all the transfers applied and can be used to save or revert to a
particular state.\
The whole point of this program is to provide functionality of performing
transfers between bags.
