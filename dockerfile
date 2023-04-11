FROM mysql:8.0

# Set the environment variables
ENV MYSQL_DATABASE mydatabase
ENV MYSQL_USER user
ENV MYSQL_PASSWORD password
ENV MYSQL_ROOT_PASSWORD password

# Copy the SQL script to create the table
COPY schema.sql /docker-entrypoint-initdb.d/

# Expose the MySQL port
EXPOSE 3306