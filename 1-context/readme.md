## Context: operation canceling and data passing ##


Often in a working project, there is a need for control over the execution of individual program branches and the ability to cancel them. Let's consider this with an example.

Suppose there is a website for searching flights, a search service, and a database with all the necessary information. A user visits the website and searches for tickets from Dubai to Tbilisi on May 8th. They click the "Search" button and initiate the process of searching and calculating the ticket prices in the service. The service goes to the database for information, but the database is overloaded, so the search takes several tens of seconds.

A few seconds later, the user realizes that they want to fly to Tbilisi on May 7th instead, as he has made plans to meet a friend. They cancel the search by clicking the "Back" button and start looking for tickets on May 7th.

What happens in the service?
```go
    import "context" 
```