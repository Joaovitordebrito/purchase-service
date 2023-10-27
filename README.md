# # purchase-service

## Wow to setup and run the project

  `git clone git@github.com:Joaovitordebrito/purchase-service.git`
  or 
  `git clone https://github.com/Joaovitordebrito/purchase-service.git`

then run the docker compose
`docker compose up --build`

## Routes
### Store a purchase transaction:
 **POST** 
 
     http://localhost:8080/purchase

 #### Request body exemple: 

      {
    	"description":"test description",
    	"purchaseAmount": 7.98
    	"transactionDate": "2023-02-03"
      }

- **description is a string**
- **purchaseAmount is a number**
- **transactionDate is a string but must be a valid date format (YYYY-MM-dd)**

**Response status code**: <span style="color:green">*201*</span>.


**Response body exemple:**

    {
    	"UUID": "db7e9990-cc3c-4c86-9105-3803b754b8c6",
    	"description": "test descrioption",
    	"transactionDate": "2023-02-03",
    	"purchaseAmount": 7.98 
    }


### Retrieve a purchase transaction in a specified country’s currency
**GET** 
 
     http://localhost:8080/converted/currency/:uuid/country

 #### Request uri exemple: 

      http://localhost:8080/converted/currency/db7e9990-cc3c-4c86-9105-3803b754b8c6/argentina

**Response status code: 201**


**Response body exemple:**

    {
		"purchaseAmount": 76.51,
		"targetCurrency": 365.5,
		"convertedAmount": 27964.41
	}

- **purchaseAmount is the value of the purchase**
- **targetCurrency is the value of the chosen currency**
- **convertedAmount is the value of the purchase converted to the chosen currency**

If the selected country does not have a currency data in the past 6 months the response will be:

    {
		"error": "the purchase cannot be converted to the target currency"
	}
