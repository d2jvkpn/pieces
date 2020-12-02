FROM base

RUN yum -y update
ADD apps /apps

# RUN yum -y groupinstall "Development Tools"
RUN yum -y install python3
RUN mkdir -p /root/.pip && cp /apps/pip.conf /root/.pip
RUN pip3 install --upgrade pip --no-cache-dir

RUN pip3 install --no-cache-dir \
    jupyter jupyterlab ipython virtualenv cpython \
    csvkit h5py PyMySQL redis \
    pyyaml toml \
    fire pinyin nose

RUN pip3 install --no-cache-dir \
    numpy pandas xlrd xlwt \
    sympy scipy statsmodels patsy \
    matplotlib seaborn pillow graphviz pydot \
    networkx

RUN pip3 install --no-cache-dir \
    BeautifulSoup4 html5lib requests flask \
    grpcio grpcio-tools pika

RUN pip3 install --no-cache-dir jieba nltk \
    xgboost scikit-learn scikit-image tensorflow

RUN rm -rf /apps
RUN yum -y autoremove; yum -y clean all
