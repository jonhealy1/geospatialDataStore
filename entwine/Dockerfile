FROM continuumio/miniconda

RUN mkdir /entwine_files
WORKDIR /entwine_files
COPY ./files /entwine_files



RUN conda install nodejs -y
RUN npm install http-server -g

ENTRYPOINT ["http-server", "/entwine_files", "-p", "8089", "--cors"]
EXPOSE 8089 8111