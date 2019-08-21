package com.geospatial.config;

import com.geospatial.entities.Client;
import com.geospatial.handlers.CustomAuthenticationSuccessHandler;
import com.geospatial.services.ClientService;
import com.geospatial.services.impl.UserDetailsServiceImpl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;


@Configuration
@EnableWebSecurity
public class WebSecurityConfig extends WebSecurityConfigurerAdapter{


    @Autowired
    private CustomAuthenticationSuccessHandler loginSuccessHandler; 

    @Autowired
    private ClientService clientService;

    @Autowired
    private UserDetailsServiceImpl userDetailsService;

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
            .csrf().disable()
            .authorizeRequests()
                .antMatchers("/static/**", "/", "/signupClient").permitAll()
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

    // @Bean
    // @Override
    // public UserDetailsService userDetailsService() {
    //     UserDetails client =
    //          User.withDefaultPasswordEncoder()
    //             .username("masterClient")
    //             .password("pass123")
    //             .roles("CLIENT")
    //             .build();
    //     return new InMemoryUserDetailsManager(client);
    // }

    @Override
    protected void configure(AuthenticationManagerBuilder auth) throws Exception {
        Client manager = new Client("manager", "pass123");
            manager.setRoles("CLIENT");
        try {
            this.clientService.registerClient(manager);
        } catch (Exception e){
            System.out.println("manager failed or already exists...");
        }
        auth.userDetailsService(this.userDetailsService);

    }
   
    @Override
    public void configure(WebSecurity web) throws Exception {
        //Web resources
     
            web.ignoring().antMatchers("/css/**");
            web.ignoring().antMatchers("/js/**");
            web.ignoring().antMatchers("/images/**");
        
       
    }


    

}

