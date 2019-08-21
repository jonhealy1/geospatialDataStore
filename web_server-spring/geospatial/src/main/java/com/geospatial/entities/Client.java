package com.geospatial.entities;

import org.springframework.data.annotation.Id;

public class Client {

    @Id
    public String id;

    public String name;

    public String password;

    public String[] roles;





    // Constructors
    public Client() {}

    public Client(String name, String password) {
        this.name = name;
        this.password = password;
    }

    // Functions
    public String getId() {
    	return this.id;
    }
    public void setId(String id) {
    	this.id = id;
    }
    public String getname() {
        return this.name;
    }
    public void setname(String name) {
        this.name = name;
    }
    public String getPassword() {
        return this.password;
    }
    public void setPassword(String password) {
        this.password = password;
    }

    public String[] getRoles() {
        return this.roles;
    }
    public void setRoles(String... roles) {
        this.roles = roles;
    }

    @Override
    public String toString() {
        return String.format(name);
    }

}