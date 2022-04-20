# Inventory Buddy

Working with many Shopify clients, sometimes there are high-demand product launches where we need to monitor stock levels. It's not always easy to keep an eye on your inventory, as Shopify does not offer this feature by default. Inventory Buddy makes sure Shopify store owners will be alerted when inventory levels drop below acceptable levels so that they can act quickly to ensure customers get the products they need.

Inventory Buddy constantly monitors your Shopify store to make sure products are well-stocked. You can specify a minimum stock level, and when inventory reaches this amount, an SMS and email will be sent to alert the store owner. It's a very simple app that solves a clear purpose: keeping online store owners aware of what is going on in their store.

## Get started

Clone or download this repository to your filesystem.

To run this project locally, you will need to configure the .env file with the follow:

Mailtrap user/pass, which you can get here: https://mailtrap.io/
Twilio API credentials: https://www.twilio.com/docs/iam/credentials/api

Additionally, you will need to enter your Shopify access token, which can be obtained here:
https://shopify.dev/apps/auth/admin-app-access-tokens

Finally, the Variant ID of the product you'd like to monitor will also be needed. This can be found by following this brief instruction: https://help.shopify.com/en/manual/products/variants/find-variant-id

Navigate to the project directory, and run the following command:
**go run main.go**
The project will now be running at http://localhost:3000/
