Domain-Driven Design pattern is the talk of the town today.

Domain-Driven Design(DDD) is an approach to software development that simplifies the complexity developers face by connecting the implementation to an evolving model.

Why DDD?
The following are the reasons to consider using DDD :

- Provide principles & patterns to solve difficult problems
- Base complex designs on a model of the domain
- Initiate a creative collaboration between technical and domain experts to iteratively refine a conceptual model that addresses domain problems.

DDD comprises of 4 Layers :

- Domain: This is where the domain and business logic of the application is defined.
- Infrastructure: This layer consists of everything that exists independently of our application: external libraries, database engines, and so on.
- Application: This layer serves as a passage between the domain and the interface layer. The sends the requests from the interface layer to the domain layer, which processes it and returns a response.
- Interface: This layer holds everything that interacts with other systems, such as web services, RMI interfaces or web applications, and batch processing frontends.

type Food struct {
	ID          uint64     `json:"id"`
	UserID      uint64     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	FoodImage   string     `json:"food_image"`
}

type User struct {
	ID        uint64     `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
}
