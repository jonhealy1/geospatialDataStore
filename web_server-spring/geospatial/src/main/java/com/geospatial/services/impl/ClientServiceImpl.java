package com.geospatial.services.impl;

import java.util.List;

import com.geospatial.dtos.ClientDTO;
import com.geospatial.entities.Client;
import com.geospatial.exceptions.UsernameExistsException;
import com.geospatial.repositories.ClientRepo;
import com.geospatial.services.ClientService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
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

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Override
    public Client getClientByUsername(String username){
        return clientRepo.findClientByUsername(username);
    }

    @Override
    public String deleteClientByUsername(String username){
        return "deleted " + username;
    }

    @Override
    @Transactional
    public Client registerClient(Client client) throws UsernameExistsException{
        Client check = clientRepo.findClientByUsername(client.getname());
        if(check != null) {
            throw new UsernameExistsException("That username is already in use");
        }
        client.setPassword(passwordEncoder.encode(client.getPassword()));
        return clientRepo.save(client);

    }
    
    @Override
    @Transactional
    public Client createClientDTO(ClientDTO clientDto) throws UsernameExistsException {
        Client check = clientRepo.findClientByUsername(clientDto.getname());
        if(check != null){
            throw new UsernameExistsException("That username is already in use");
        }
        Client newClient = new Client(clientDto.getname(), passwordEncoder.encode(clientDto.getPassword()));
        newClient.setRoles("CLIENT");
        return clientRepo.save(newClient);
    }
}