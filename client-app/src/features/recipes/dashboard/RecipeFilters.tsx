import { Fragment } from "react";
import { Menu, Header } from "semantic-ui-react";
import React from "react";

const RecipeFilters = () => (
    <Fragment>
        <Menu vertical size={'large'} style={{width: '100%', marginTop: 50}}>
            <Header icon={'filter'} attached color={'teal'} content={'Filters'}/>
            <Menu.Item color={'blue'} name={'all'} content={'All Recipes'}/>
        </Menu>
    </Fragment>
)

export default RecipeFilters;