package com.geospatial.web;

import java.util.List;

import javax.validation.Valid;

import com.geospatial.dtos.ClientDTO;
import com.geospatial.entities.Client;
import com.geospatial.exceptions.UsernameExistsException;
import com.geospatial.repositories.ClientRepo;
import com.geospatial.services.ClientService;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.ui.ModelMap;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
public class ClientController {

    private final Logger logger = LoggerFactory.getLogger(MainController.class);

    @Autowired
    private ClientService clientService;

    @Autowired
    private ClientRepo clientRepo;

    public ClientController() {
        logger.debug("Client Controller Initialized");
    }

    @RequestMapping(value="/clients", method = RequestMethod.POST, produces = "application/json")
    @ResponseBody
    public List<Client> getClients(Model model) {
        return this.clientService.getClients();
    }

    @RequestMapping("/deleteClients")
    @ResponseBody
    public String deleteClients(){
        clientRepo.deleteClients();
        return "success";
    }

    @RequestMapping(value="/signupClient", method = RequestMethod.POST)
    public ModelMap createClient(@ModelAttribute("clientDto") @Valid ClientDTO clientDto, BindingResult result) throws UsernameExistsException{
        ModelMap m = new ModelMap();
        try{
            clientService.createClientDTO(clientDto);
        } catch (UsernameExistsException e){
            result.rejectValue("name", "clientDto", "Username already exists...");
        }
        if(result.hasErrors()){
            m.addAttribute("clientDto", clientDto);
            return m;
        }  
        // System.out.println("FIRE");
        // System.out.println(clientDto.getname());
        // System.out.println(clientDto.getPassword());
        // System.out.println(clientDto.getMatchPassword());
              
       
   
    return new ModelMap("clientDto", new ClientDTO());
    }









}