# Restaurant-Management
This is a simple REST API for managing restaurant menus, food orders,order items, invoices, tables and users. It includes basic CRUD operations and uses password hashing to ensure secure authentication. Tokens are generated upon login for additional security.


## Technologies Used
MongoDB


### API Endpoints
- Food:
   - GET("/foods", controller.GetFoods())
   - GET("/foods/:food_id", controller.GetFood())
	- POST("/foods", controller.CreateFood())
	- PATCH("/foods/food_id", controller.UpdateFood())
- Invoice:
  - GET("/invoices", controller.GetInvoices())
	- GET("/invoices/invoice_id", controller.GetInvoice())
	- POST("/invoices", controller.CreateInvoice())
	- PATCH("/invoices/Invoice_id", controller.UpdateInvoice())
- Menu:
  - GET("/menus", controller.GetMenus())
	- GET("/menus/menu_id", controller.GetMenu())
	- POST("/menus", controller.CreateMenu())
	- PATCH("/menus/menu_id", controller.UpdateMenu())
- Order:
  - GET("/Oder", controller.GetOrders())
	- GET("/Oder/Oder_id", controller.GetOrder())
	- POST("/Oder", controller.CreateOrder())
	- PATCH("/Oder/Oder_id", controller.UpdateOrder())
 - Order Items:
  - GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/order_id", controller.GetOrderItemByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("orderItems/orderItem_id", controller.UpdateOrderItem())
