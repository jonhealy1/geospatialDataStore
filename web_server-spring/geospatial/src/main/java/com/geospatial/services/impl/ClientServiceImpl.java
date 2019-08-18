package com.geospatial.services.impl;

import java.util.List;

import com.geospatial.dtos.ClientDTO;
import com.geospatial.entities.Client;
import com.geospatial.repositories.ClientRepo;
import com.geospatial.services.ClientService;

import org.assertj.core.util.Arrays;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class ClientServiceImpl implements ClientService {
    
    @Autowired
    private ClientRepo clientRepo;

    @Override
    public List<Client> getClients() {
        return clientRepo.getClients();
    }

    @Override
    public Client getClientByUsername(String username){
        return clientRepo.getClientByUsername(username);
    }
    
    @Override
    @Transactional
    public Client createClient(ClientDTO clientDto) {
        Client newClient = new Client(clientDto.getUsername(), clientDto.getPassword());
        newClient.setRoles(Arrays.asList("ROLE_CLIENT"));
        return clientRepo.save(newClient);
    }
}