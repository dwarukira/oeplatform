import React from "react";
import logo from "./logo.svg";
import "./App.css";
import "antd/dist/antd.css";
import styled from "styled-components";
import { Router, Link } from "@reach/router";

import DashboardLayout from "./componets/layout";
import UserList from "./pages/Users/list";
import UserEdit from "./pages/Users/edit";
import RoleList from "./pages/Users/roles";
import RoleEdit from "./pages/Users/roles_edit";
import SellerList from "./pages/Directory/sellers";
import SellerDetails from "./pages/Directory/seller_details";
import RoleCreate from "./pages/Users/roles_new";
import routes from "./routes";
import NotFound from "./componets/NotFound";

const Container = styled.div`
  height: 100vh;
  width: 100vw;
`;

const App: React.FC = () => {
  return (
    <Container>
      <DashboardLayout>
        <Router>
          <NotFound default />
          <UserList path="/users/list" />
          <RoleList path="/users/roles" />
          <RoleEdit path="/users/role/:id" />
          <UserEdit path="/user/:id" />
          <SellerList path={routes.directory_sellers} />
          <SellerDetails path={routes.directory_seller} />
          <RoleCreate path={routes.roles_create} />
        </Router>
      </DashboardLayout>
    </Container>
  );
};

export default App;
