Convention 2:
Always use one repo per microservice, unless you decide to have a mono repo (which I recommend for all small teams).

Best practice, deploy before you code.

When working on large scale projects, one of the most important things that I've discoved in my 12+ years of software development is the importance of consistent conventions. 
Imagine starting out at company XYZ and being tasked with familiarizing yourself with microservice A that lives in repo A. You spend weeks getting your environment set up; you 
even mangage fix a broken unit test that you noticed was broken while carefully following the README docs. Finally, you feel convertable navigating repo A, and you even feel
your first stroke of pride seeing your first PR get merged by your new team. But as you bask in your new found moment, your team notifies you that they need you to start working on
microservice D. You dread the feeling of being a rookie again. Imagine if it didn't have to be so hard.

Solution:
I love boring conventions -- it makes me excited. It allows you easily naviate from repo A to repo B, etc., and only have to learn your company's business logic for each microserivce, but you
only have to learn the conventions once.

Convention 1:
Every backend developer, irregardless of the programming language has worked seen some version of the following problem: take some input data, transform it, and do something. If you
are a functional programmer, you probably see every problem as such. Nevertheless, the I want to propose a grand idea. Always, and I say always create a DTO model for every model 
that you have.

Pros:
It becomes very easy to separate the responsibilities of modelling input data. In my experience, responsibilities of capturing data can be delegated to more junior developer or developers that
are new to the microservice. Developers, more familiar with the microservices business logic

In this project, we will display three news articles to our users.

If you are following along, you can sign up at https://www.marketaux.com/ and get a free api key in 30 seconds.

First create a models directory.

In it, we create four files: 
`market_aux_data_dto.go`
`market_aux_meta_dto.go`
`market_aux_entities_dto.go`
`market_aux_request_dto.go`

Here are three coding standards that I like to enforce:
1. Only one model per file: I find it very nice to be able to browse the projects models directory and see all the names of each model.
2. Name the file as the exact name as the model.
3. Use name prefixes to group similar models so that they are grouped together. E.g. all models requested to requesting data from marketaux.com start with market_aux_XYZ.go.

We then copy for the dto files and remove the `_dto` from the filename to be our actual models:
`market_aux_data.go`
`market_aux_meta.go`
`market_aux_entities.go`
`market_aux_request.go`

The next convention that we will adopt is that each DTO model will have a ToModel() function that will return a pointer to the acutal model.