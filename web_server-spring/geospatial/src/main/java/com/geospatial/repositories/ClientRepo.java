package com.geospatial.repositories;


import java.util.List;

import com.geospatial.entities.Client;

import org.springframework.data.mongodb.repository.DeleteQuery;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

public interface ClientRepo extends MongoRepository<Client, String> {

    @Query("{}")
    List<Client> getClients();

    @Query("{ 'username ': ?0 }")
    Client findClientByUsername(String username);

    @DeleteQuery("{ 'username ': ?0 }")
    Client deleteClientByUsername(String username);



    
}