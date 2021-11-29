# Getting Started with Gauge

This is an example markdown specification file.
Every heading in this file is a scenario.
Every bulleted point is a step.

To execute this specification, use
	npm test

This is a context step that runs before every scenario
* Open shopping-cart application

## Add product to basket 
* Add first product to basket
* Go to "Basket" page
* See the product

## Create new Product
* Go to "Create" page
* Write to "name" area "test product"
* Write to "price" area "1"
* Write to "image" area "https://bireysel.turktelekom.com.tr/cihazlar/publishingimages/telefon/iphone13-blue.jpg"
* Create the product

## Update product
* Add first product to basket
* Go to "Basket" page
* "Increase" "2" times countity of product
* "Decrease" "1" times countity of product
* Delete product