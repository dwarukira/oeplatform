import React, { Children } from "react";
import { Layout, Menu, Icon } from "antd";
import styled from "styled-components";
import { Link, RouteComponentProps, Location } from "@reach/router";
import Logo from "../Logo";

const { Header, Sider, Content } = Layout;
const { SubMenu } = Menu;

const Sub = styled.div`
  height: 16px;
  width: 64px;
  color: rgba(255, 255, 255, 0.5);
  font-family: Roboto;
  font-size: 14px;
  font-weight: bold;
  letter-spacing: 0.58px;
  line-height: 16px;
  margin-left: 30px;
  margin-bottom: 20px;
  margin-top: 20px;
`;


interface IProps extends RouteComponentProps {

}

class DashboardLayout extends React.Component<IProps> {
  state = {
    collapsed: false
  };

  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed
    });
  };

  render() {
    return (
      <Location>
      {({location}) => {

        return (
      <Layout style={{ minHeight: "100vh" }}>
        <Sider trigger={null} collapsible collapsed={this.state.collapsed}>
          <div className="logo">
            <Logo />
            </div>

          <Menu 
            theme="dark" 
            mode="vertical" 
            defaultSelectedKeys={[location.pathname]}
          >
            <Sub> Analytics </Sub>
            <Menu.Item key="1">
              <Icon type="appstore" />
              <span>Dashboard</span>
            </Menu.Item>
            <Menu.Item key="2">
              <Icon type="rise" />
              <span>Reports</span>
            </Menu.Item>

            <Sub> Platform </Sub>
            <Menu.Item key="3">
              <Icon type="profile" />
              <span>Content</span>
            </Menu.Item>
            <Menu.Item key="4">
              <Icon type="video-camera" />
              <span>Inventory</span>
            </Menu.Item>
            <Menu.Item key="5">
              <Icon type="shopping-cart" />
              <span>Orders</span>
            </Menu.Item>

      
            <SubMenu
              key="6"
              title={
                <span>
                  <Icon type="file-search" />
                  <span>Directory</span>
                </span>
              }
            >
              <Menu.Item key="/directory/customers"> <Link to="/directory/customers"> Customers </Link> </Menu.Item>
             <Menu.Item key="/directory/sellers"> <Link to="/directory/sellers" > Sellers </Link></Menu.Item> 
            
              
            </SubMenu>

            <Sub> Finance </Sub>

            <Menu.Item key="8">
              <Icon type="money-collect" />
              <span>Payments</span>
            </Menu.Item>

            <Menu.Item key="9">
              <Icon type="wallet" />
              <span>Wallet</span>
            </Menu.Item>

            <Sub> Account </Sub>

            <Menu.Item key="10">
              <Icon type="notification" />
              <span>Notifications</span>
            </Menu.Item>

            <SubMenu
              key="11"
              title={
                <span>
                  <Icon type="usergroup-add" />
                  <span>Users</span>
                </span>
              }
            >
             <Menu.Item key="/users/list"> <Link to="/users/list" >Add/Edit </Link></Menu.Item> 
              <Menu.Item key="/users/roles"> <Link to="/users/roles"> Roles </Link> </Menu.Item>
              <Menu.Item key="user_ac">My Account</Menu.Item>
            </SubMenu>

            <Menu.Item key="12">
              <Icon type="setting" />
              <span>Settings</span>
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout>
          <Header style={{ background: "#fff", padding: 0 }}>
            <Icon
              className="trigger"
              type={this.state.collapsed ? "menu-unfold" : "menu-fold"}
              onClick={this.toggle}
            />
          </Header>
          <Content
            style={{
              margin: "24px 16px",
              padding: 24,
              background: "#fff",
              minHeight: 280
            }}
          >
            { this.props.children }
          </Content>
        </Layout>
      </Layout>
        )
      }}
          </Location>

    );
  }
}

export default DashboardLayout;
