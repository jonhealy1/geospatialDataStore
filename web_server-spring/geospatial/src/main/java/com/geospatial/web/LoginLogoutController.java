package com.geospatial.web;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.geospatial.dtos.ClientDTO;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.autoconfigure.web.servlet.WebMvcProperties.View;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.logout.SecurityContextLogoutHandler;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
public class LoginLogoutController {

    private final Logger logger = LoggerFactory.getLogger(LoginLogoutController.class);

    public LoginLogoutController() {
        logger.debug("LoginLogout controller initialized");
    }


    @RequestMapping(value={"/login", "/", ""}, method = RequestMethod.GET)
    public String login(Model model) {
        ClientDTO clientDto = new ClientDTO();
        model.addAttribute("clientDto", clientDto);
        return "login";
    }

    @RequestMapping(value="/logout", method = RequestMethod.GET)
    public String logout(HttpServletRequest request, HttpServletResponse response) {
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        if (auth != null) {
            new SecurityContextLogoutHandler().logout(request, response, auth);
        }        
        return "redirect:/login?logout";
    }

        
    @RequestMapping(value="/signupClient", method = RequestMethod.GET)
    public String signupClient(View view) {
        return "login";
    }





    

}