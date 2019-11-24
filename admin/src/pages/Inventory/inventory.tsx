import React, { FunctionComponent } from "react";

import { Table } from "antd";
import { RouteComponentProps, Link } from "@reach/router";
import { useQuery } from "@apollo/react-hooks";
import { GET_USERS } from "../../gql/queries";
import { Title } from "../../componets/Heading";


function truncate(str: string) {
  return str.length > 10 ? str.substring(0, 7) + "..." : str;
}

const columns = [
  {
    title: "ID",
    dataIndex: "id",
    render: (text: string) => <Link to={`/user/${text}`}>{truncate(text)}</Link>
  },

  {
    title: "Username",
    dataIndex: "name"
  },

  {
    title: "Name",
    dataIndex: "name"
  },

  {
    title: "Email",
    dataIndex: "email"
  },
  {
    title: "Role",
    dataIndex: "role"
  },
  {
    title: "Status",
    dataIndex: "status"
  },
  {
    title: "Last Login",
    dataIndex: "last_login"
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

const Inventory: FunctionComponent<RouteComponentProps> = () => {
  const { loading, error, data } = useQuery(GET_USERS);

  const hasData = data !== undefined && !error;
  function createColumnData() {
    return data.users.list.map((item: any) => ({
      key: item.id,
      id: item.id,
      name: item.name,
      email: item.email,
      last_login: new Date(item.lastLogin).toLocaleString()
    }));
  }

  return (
    <>
      <Title> Listing all user accounts </Title>
      <Table
        loading={loading}
        rowSelection={rowSelection}
        columns={columns}
        dataSource={hasData ? createColumnData() : null}
      />
    </>
  );
};

export default Inventory;
