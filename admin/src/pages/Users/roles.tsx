import React, { FunctionComponent } from "react";

import { Table, Row, Col, Button, Icon } from "antd";
import { RouteComponentProps, Link } from "@reach/router";
import { useQuery } from "@apollo/react-hooks";
import { GET_ROLES } from "../../gql/queries";
import { Title } from "../../componets/Heading";
import routes from "../../routes";
import styled from "styled-components";
import { PaddingTop } from "../../componets/Padding";


const AddButton = styled(Button)`
	box-sizing: border-box;
	height: 36px;
	width: 100px;
	border: 1px solid #D8DCE6;
	border-radius: 4px;
	background-color: #FFFFFF;
`;


const columns = [
  {
    title: "ID",
    dataIndex: "id",
    render: (text: React.ReactNode) => (
      <Link to={`/users/role/${text}`}>{text}</Link>
    )
  },
  {
    title: "Role",
    dataIndex: "role"
  },

  {
    title: "Last modified on",
    dataIndex: "updateAt"
  }
];

// rowSelection object indicates the need for row selection
const rowSelection = {
  onChange: (selectedRowKeys: any, selectedRows: any) => {
    console.log(
      `selectedRowKeys: ${selectedRowKeys}`,
      "selectedRows: ",
      selectedRows
    );
  },
  getCheckboxProps: (record: { name: string }) => ({
    disabled: record.name === "Disabled User", // Column configuration not to be checked
    name: record.name
  })
};

const RoleList: FunctionComponent<RouteComponentProps> = () => {
  const { loading, error, data } = useQuery(GET_ROLES);

  const hasData = data !== undefined;
  function createColumnData() {
    return data.roles.list.map((item: any) => ({
      key: item.id,
      id: item.id,
      role: item.name,
      updateAt: new Date(item.updatedAt).toLocaleString()
    }));
  }

  if (error) {
    return <p> XX </p>;
  }

  return (
    <>
      <PaddingTop>
        <Row>
          <Col sm={6}  span={12}>
            <Title> Listing all user roles </Title>
          </Col>

          <Col sm={6}  span={12}>
            <Link to={routes.roles_create}> <AddButton> 
              <Icon type="plus" />
              Add Role 
              </AddButton> </Link>
          </Col>
        </Row>
      </PaddingTop>
      <Table
        loading={loading}
        rowSelection={rowSelection}
        columns={columns}
        dataSource={hasData ? createColumnData() : null}
      />
    </>
  );
};

export default RoleList;
