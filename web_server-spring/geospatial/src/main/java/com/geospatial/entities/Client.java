package com.geospatial.entities;

import java.util.List;

import org.springframework.data.annotation.Id;

public class Client {

    @Id
    public String id;

    public String username;
    public String password;
    private List<Object> roles;



    // Constructors
    public Client() {}

    public Client(String username, String password) {
        this.username = username;
        this.password = password;
    }

    // Functions
    public String getId() {
    	return this.id;
    }
    public void setId(String id) {
    	this.id = id;
    }
    public String getUsername() {
        return this.username;
    }
    public void setUsername(String username) {
        this.username = username;
    }
    public String getPassword() {
        return this.password;
    }
    public void setPassword(String password) {
        this.password = password;
    }
    public List<Object> getRoles() {
        return this.roles;
    }
    public void setRoles(List<Object> roles) {
        this.roles = roles;
    }

    @Override
    public String toString() {
        return String.format(username);
    }

}