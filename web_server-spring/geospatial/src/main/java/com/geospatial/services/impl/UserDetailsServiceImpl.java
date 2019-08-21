package com.geospatial.services.impl;

import com.geospatial.entities.Client;
import com.geospatial.repositories.ClientRepo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;


@Service
@Transactional
public class UserDetailsServiceImpl implements UserDetailsService{

    @Autowired
    private ClientRepo clientRepo;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        

        Client client = this.clientRepo.findClientByUsername(username);
        System.out.println(username);
        System.out.println(client);

        if(client!=null){
            return User.withDefaultPasswordEncoder()
                        .username(client.getname())
                        .password(client.getPassword())
                        .roles(client.getRoles())
                        .build();
        } else{
            return null;
            
        }

    }

}