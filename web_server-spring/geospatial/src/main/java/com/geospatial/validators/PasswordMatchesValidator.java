package com.geospatial.validators;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;

import com.geospatial.dtos.ClientDTO;


public class PasswordMatchesValidator implements ConstraintValidator<PasswordMatches, Object> {

    @Override
    public void initialize(PasswordMatches constaintAnnotation) {}

    @Override
    public boolean isValid(Object value, ConstraintValidatorContext context) {
        ClientDTO client = (ClientDTO) value;
        return client.getPassword().equals(client.getMatchPassword());
    }

}