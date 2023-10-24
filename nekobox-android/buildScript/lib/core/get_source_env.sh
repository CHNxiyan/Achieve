if [ ! -z $ENV_NB4A ]; then
  export COMMIT_SING_BOX_EXTRA="046c0e406ed4719608fb38fed90e3b1e229163f9"
fi

if [ ! -z $ENV_SING_BOX_EXTRA ]; then
  source libs/get_source_env.sh
  # export COMMIT_SING_BOX="91495e813068294aae506fdd769437c41dd8d3a3"
fi
