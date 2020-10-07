# First Iteration

### Assumptions

1. Unit price is updated no more than once a week
2. Categories don't change

---

```json
Item -> {
    Id,
    CategoryName,
    CategoryDescription,
    CategoryId, // PARTITION?
    UnitPrice,

    BuyCount,
    ViewCount,
    VoteCount,
    AddToCartCount,
    ReleaseDate // PARTITION!
}

Category -> {
}

Orders -> {
    UserId,
    Id,
    Items: [
        {
            UnitPrice,
            Quantity,
            Item: {
                CategoryName
                CategoryDescription
                Name
            }
        }
    ],
    Address: {
        City,
        State,
        PostalCode,
        Address,
        Country,

    }
    HasBeenShipped
    // Total ?,
}

Cart {
    UserId,
    Id,
    Items: [
        {
            Quantity,
            UnitPrice, // updated by UDF
            Item: {
                CategoryName
                CategoryDescription
                Name
            }
        }
    ],
}

User {
    Id
}
```

- Retrieve list of categories: ~500 requests/second
    -> SELECT from Category
- Filter movies by category: ~200 requests/second
    -> SELECT from Items WHERE CategoryID = {CAT_ID}
- Retrieve list of top 10 movies by popularity (home page): ~500 requests/second
    -> SELECT TOP 10 from Items ORDER BY BuyCount DESC
- *L* Retrieve orders with details showing products with quantities: ~10 requests/second
    -> SELECT * FROM Orders WHERE Id='{SOME_ID}'
- *L* Add a movie to the shopping cart: ~5 requests/second
    -> INSERT IF NOT EXISTS INTO Cart (UserId="abe" ....., Items=[])

- *L* Complete a purchase transaction: ~2 requests/second
    -> ```
        order = Order()
        cart = getCart()
        order.items = cart.items
        db.saveOrder(order)

- *M* Retrieve list of last 10 movies by release date: ~75 requests/second
    -> SELECT TOP 10 FROM Item ORDER BY ReleaseDate ASC ?

- *M* Retrieve details for a specific movie: ~160 requests/second
- *S* Page through lists of movies: ~30 requests/second