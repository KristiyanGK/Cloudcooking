import { Fragment, useContext, useState, useEffect } from "react";
import { Menu, Header, Form, Button } from "semantic-ui-react";
import { Form as FinalForm, Field } from 'react-final-form';
import React from "react";
import { observer } from "mobx-react-lite";

const RecipeFilters = () => {
    const handleFinalFormSubmit = (values: any) => {;
        console.log(values);
    };

    return (
    <Fragment>
        <Menu vertical size={'large'} style={{width: '100%', marginTop: 50}}>
            <Header icon={'filter'} attached color={'teal'} content={'Filters'}/>
            <Menu.Item color={'blue'} name={'all'} content={'All Recipes'}/>
            <FinalForm 
            onSubmit={handleFinalFormSubmit}
            render={({ handleSubmit }) => (
                <Form onSubmit={handleSubmit}>
                    <Button
                        content="filter"
                    />
                </Form>
            )}/>
            
        </Menu>
    </Fragment>
    )
}

export default observer(RecipeFilters);