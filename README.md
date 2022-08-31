## Time.Now

This is my first ever master project. It is a collection of watches, so I named it as 'Time.Now'. There is a reason behind this name. That reason we all Gophers should knew. So in my point of view it is a right name for this project.

In this I have done whole backend works, that means Rest APIs for every endpoints. There is having a question that why I done this project, If you want to know that please keep read the following.

Before doing this project, I thought which project I will choose. That time I asked to some of my friends, then they suggest me if you can make an E-commerce website, then go with that. And while you are doing an E-commerce you can learn a lot. Basically I want to learn deeply what I am learning, it is one of my hobby. So I went with the E-commerce.

The very beginning of my project I understood that making an E-commerce website is not easy, it should take a lot of time. But I didn't consider anything, I only concentrate on my E-commerce project. For doing this project I didn't follow any project tutorials, and I have done my project alone, So this is my confidence.

This project is completely dockerized. If you want to run my project.. execute the following command (need to installed docker in your machine) :

      docker compose up

If you execute this command, it will create 2 containers. The very first postgres container will create. After that it will do database migration automatically in postgres container. After that App container will create. Then the App will connect with the database that is postgres container.

If you want to see the logs, execute the following commands:

      docker-compose logs

Talking about the features of my project, first I wanna say about the architecture which I used. Code Architecture is the main thing in a big project, because we all know, in a company we don't make a project alone. So for a teamwork we should use a better architecture. And here I followed the MVC architecture.

The second one, I used Golint package. Talking about Golint, it is a linter maintained by the Go developers. It is intended to enforce the coding conventions described in the Effective Go and CodeReviewComments. These same conventions are used in the open-source Go project and at Google also.

In my project I have used PostgreSQL as database. PostgreSQL is an advanced open-source object-relational database system which applies SQL language. As an RDBMS it knows how to handle concurrency and very diverse workloads.

Next one I have used Goose package. It is a database migration tool. It manages database schema by creating incremental SQL changes or Go functions. 
And I used more features, I will explain those in detail later.

The first page or landing page of the project is a products listing page. In that page users can view limited amount of products, because I have used pagination when listing products. In this page users can select a product and he should select it's color as well. After that for buying that product, user have 2 options, that means either can add that product to cart or can buy that product by instant buy method. Both these 2 methods have 2 differnet type of payments, Cash on delivery and Banking. 

To be continued..
