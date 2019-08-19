package com.geospatial.dtos;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;

// import com.geospatial.validators.ValidPassword;

public class ClientDTO{

    @NotNull
    @NotEmpty
    private String username;




    @NotNull
    @NotEmpty
    // @ValidPassword
    private String password;



    // private String matchPassword;


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
    // public String getMatchPassword() {
    //     return this.matchPassword;
    // }
    // public void setMatchPassword(String matchPassword) {
    //     this.matchPassword = matchPassword;
    // }

}