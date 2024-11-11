# grpc-enricher

The purpose of this repository is to demonstrate how to implement GraphQL-like dynamic enrichment for gRPC-based services, using go.

In a typical GraphQL server, we encounter two types of enrichment:
- Field Enrichment: This allows us to specify which fields we want to retrieve when querying a given entity.
- Relation Enrichment: If an entity contains a nested entity (i.e., a foreign ID), we can specify whether to retrieve that nested entity, while simultaneously handling field enrichment.

## Project description

In this project, I will illustrate the techniques using a simple blog example. The blog has the following relationships:
It allows the creation of posts, categories, and tags. Each post belongs to a category and may contain one or more tags. Each category may also contain one or more tags. While this structure may not be entirely realistic, it effectively illustrates how nested enrichment functions.

The only available gRPC service endpoint is the ListPost method. This endpoint is served by the grpc-server application, which is called by a test application (grpc-client).

In this demonstration, I use dummy data, which can be found in the data.go file within each domain module.

## Project Structure

If you're interested in understanding the rationale behind the folder and module structure implementation, I recommend reviewing the following repository: https://github.com/kazmerdome/best-ever-golang-starter.

Domains
- category
- post
- tag

Available Providers
- enricher: Contains the logic for field and relation enrichment within each domain.
- dataloader: A utility developed by Facebook that collects unique requests and forwards them in batches to the repository layer.
- repository: Responsible for stateful operations. Currently operates on dummy data.
- controller: Implements the gRPC server service.
- module: Assists with dependency exporting and importing.
- grpc folders in domain modules: contain the proto files and generated codes.

Entry Points:
- grpc-server
- grpc-client

## Run the applications
1. run grpc-server
2. run grpc-client
3. check the logs of both services

As soon as you run the client application, you can see in the server logs exactly how many requests are made to the repository layer (yellow logs) and how many requests are made to the dataloaders (blue logs). This clearly demonstrates how effective the dataloader technique is, as without it, fetching 10 posts would result in 10 category requests, 10 tag requests, and then an additional 10 tag detail requests, totaling 30+1 requests to the repository layer. In our case, this is reduced to only 3+1 requests.


## Enricher

### Key Features

1. Field and Self-Enrichment:
   - The Enrich method supports selective field enrichment by checking the <entity>EnrichParams, allowing clients to retrieve only the fields they need. This minimizes data transfer and improves performance by only loading fields and relations that the client explicitly requests.
   - Self-enrichment is managed by loading core data fields of the given entity (e.g., id, title, tags, category) using the <entity>Dataloader to fetch from the repository layer.

2. Relation Enrichment with Concurrent Processing:
   - The enricher also supports relation enrichment (e.g., loading related tags and categories). Using Go routines and channels, the method processes category and tag enrichment in parallel, leveraging asynchronous calls to improve response time.
   - Relation enrichment parameters (CategoryEnrichParams and TagEnrichParams) allow clients to specify nested fields and relationships, enabling more granular control over returned data.

3. Flexible Enrichment Parameters:
   - The enrichment parameters (EnrichAllFields and EnrichAllRelations) allow toggling between different enrichment levels dynamically, providing a flexible API for different data retrieval scenarios.

This approach allows a gRPC service to offer flexible data retrieval similar to GraphQL, making it easier to fetch tailored data efficiently. The method leverages Go’s concurrency model for efficient nested relation loading, creating a powerful enrichment solution for complex, relational data in microservices.


## Dataloader

The dataloader pattern, popularized by Facebook's GraphQL, is a technique designed to efficiently batch and cache requests for related data, minimizing redundant data fetching and improving performance, especially in data-intensive applications. This Go implementation leverages the dataloader package to handle batches of data requests for tags, allowing for optimal interaction with the database or repository layer.

### Key Components in dataloader (f.e.tagDataloader)

1. Batched Loading
   - The dataloader.NewBatchedLoader method sets up batched data loading for both single and multiple items.
   - The ItemLoader and ItemsLoader methods are responsible for individual and bulk tag fetching requests, respectively. Each loader creates a "thunk" (a deferred execution function), which is eventually triggered when the batch is ready to execute.

2. Batch Execution
   - The batchItemLoader and batchItemsLoader methods convert requested keys into UUIDs and prepare placeholder maps (bucket) for batch results.
   - After receiving IDs, the loaders call GetManyByIds from the repository layer, fetching all items in a single request and distributing results back to the correct placeholders.
   - These methods use Go channels and async functions to process batched results and handle errors, enabling concurrency.
  
3. NoCache
   - This implementation of the dataloader disables caching with dataloader.NoCache{}, which can be useful in scenarios where data consistency is critical, and caching would introduce stale data. However, caching can be enabled based on use case requirements.

This setup ensures that all individual or bulk requests for tags are aggregated into batched repository calls, reducing the number of total database or API requests and improving efficiency by consolidating related queries. The dataloader pattern in this context makes it well-suited for enriching complex, relational data structures by handling dependencies efficiently.
