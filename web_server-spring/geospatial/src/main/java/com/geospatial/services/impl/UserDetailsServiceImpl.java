package com.geospatial.services.impl;

import com.geospatial.entities.Client;
import com.geospatial.repositories.ClientRepo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.User.UserBuilder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;


@Service
@Transactional
public class UserDetailsServiceImpl implements UserDetailsService{



    @Autowired
    private ClientRepo clientRepo;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        

        Client client = this.clientRepo.findClientByUsername(username);
        System.out.println(username);
        System.out.println(client);

        UserBuilder builder = null;

        if(client!=null){

            builder = org.springframework.security.core.userdetails.User.withUsername(username);
            builder.password(client.getPassword());
            builder.roles(client.getRoles());

            // return User.
            // .username(client.getname())
            //             .password(client.getPassword())
            //             .roles(client.getRoles())
            //             .build();
        } else{
            throw new UsernameNotFoundException("User not found.");
            
        }
        return builder.build();

    }

}