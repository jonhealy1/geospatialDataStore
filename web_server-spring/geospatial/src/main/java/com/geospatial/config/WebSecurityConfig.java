package com.geospatial.config;

import com.geospatial.handlers.CustomAuthenticationSuccessHandler;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.provisioning.InMemoryUserDetailsManager;


@Configuration
@EnableWebSecurity
public class WebSecurityConfig extends WebSecurityConfigurerAdapter{


    @Autowired
    private CustomAuthenticationSuccessHandler loginSuccessHandler; 


    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
            .csrf().disable()
            .authorizeRequests()
                .antMatchers("/static/**", "/").permitAll()
                .anyRequest().authenticated()
            .and()
                .formLogin().successHandler(this.loginSuccessHandler)
                .loginPage("/login").permitAll()
            .and()
                .logout().permitAll();
                // .logoutUrl("/logout")
                // .logoutSuccessUrl("/login")
                // .invalidateHttpSession(true)
                // .deleteCookies("JSESSIONID")
                // .logoutSuccessHandler(this.logoutSuccessHandler);
    }

    @Bean
    @Override
    public UserDetailsService userDetailsService() {
        UserDetails client =
             User.withDefaultPasswordEncoder()
                .username("masterClient")
                .password("pass123")
                .roles("CLIENT")
                .build();
        return new InMemoryUserDetailsManager(client);
    }
   
    @Override
    public void configure(WebSecurity web) throws Exception {
        //Web resources
     
            web.ignoring().antMatchers("/css/**");
            web.ignoring().antMatchers("/js/**");
            web.ignoring().antMatchers("/images/**");
        
       
    }


    

}

