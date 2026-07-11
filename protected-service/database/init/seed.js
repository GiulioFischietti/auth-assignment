db = db.getSiblingDB("orders_db");


db.orders.insertMany([

    {
        _id: "ord001",

        customer: {
            id: 1,
            username: "mario"
        },

        payment: {
            status: "paid"
        },

        order_status: "delivered",

        items: [
            {
                product: "Laptop",
                quantity: 1,
                price: 1200
            }
        ]
    },


    {
        _id: "ord002",

        customer: {
            id: 1,
            username: "mario"
        },

        payment: {
            status: "paid"
        },

        order_status: "processing",

        items: [
            {
                product: "Mouse",
                quantity: 2,
                price: 50
            }
        ]
    },


    {
        _id: "ord003",

        customer: {
            id: 2,
            username: "luigi"
        },

        payment: {
            status: "blocked"
        },

        order_status: "cancelled",

        items: [
            {
                product: "Keyboard",
                quantity: 1,
                price: 100
            }
        ]
    }

]);