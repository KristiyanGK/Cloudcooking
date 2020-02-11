import React from "react";
import { Menu, Container, Button } from "semantic-ui-react";
import { observer } from "mobx-react-lite";
import { NavLink } from "react-router-dom";

const NavBar: React.FC = () => {
  return (
    <Menu fixed="top" inverted>
      <Container>
        <Menu.Item header exact as={NavLink} to="/">
          <img
            src="/assets/logo.png"
            alt="logo"
            style={{ marginRight: "10px" }}
          />
          Cloud Cooking
        </Menu.Item>
        <Menu.Item name="Recipes" as={NavLink} to="/recipes" />
        <Menu.Item>
          <Button
            as={NavLink}
            to="/createRecipe"
            positive
            content="Create Recipe"
          />
        </Menu.Item>
      </Container>
    </Menu>
  );
};

export default observer(NavBar);
