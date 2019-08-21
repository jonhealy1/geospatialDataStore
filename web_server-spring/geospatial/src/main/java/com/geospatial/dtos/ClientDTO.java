package com.geospatial.dtos;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;

import com.geospatial.validators.ValidPassword;



public class ClientDTO{

    @NotNull
    @NotEmpty
    private String name;

    @NotNull
    @NotEmpty
    @ValidPassword
    private String password;

    @NotNull
    @NotEmpty
    @ValidPassword
    private String matchPassword;


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
    public String getMatchPassword() {
        return this.matchPassword;
    }
    public void setMatchPassword(String matchPassword) {
        this.matchPassword = matchPassword;
    }

}