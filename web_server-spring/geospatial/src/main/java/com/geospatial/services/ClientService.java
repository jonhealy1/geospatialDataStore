package com.geospatial.services;

import java.util.List;

import com.geospatial.dtos.ClientDTO;
import com.geospatial.entities.Client;
import com.geospatial.exceptions.UsernameExistsException;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public interface ClientService {

    public List<Client> getClients();

    public Client getClientByUsername(String username);

    @Transactional
    public Client registerClient(Client client) throws UsernameExistsException;

    @Transactional
    public Client createClientDTO(ClientDTO clientDto); 

    @Transactional
    public String deleteClientByUsername(String username);


}